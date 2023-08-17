package apiHandlers

import (
	"fmt"
	"net/http"
	"pomo/internal/config"
	modelsInput "pomo/internal/models/input"
	modelsResponse "pomo/internal/models/response"
	"pomo/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Login the user and return the token
func (h *Handler) LoginHandler(ctx *gin.Context) {
	var input modelsInput.LoginInput
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := utils.GoValidator(&input)

	if errCount > 0 {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, errResponse)
		return
	}
	user, errLogin := h.service.Login(&input)

	switch errLogin {

	case http.StatusNotFound:
		utils.APIResponse(ctx, "User account is not registered", http.StatusNotFound, nil)
		return

	case http.StatusUnauthorized:
		utils.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, nil)
		return

	case http.StatusAccepted:
		// Generate Tokens
		accessToken, err := utils.CreateToken(config.Config.AccessTokenExpiresIn, user.ID, config.Config.AccessTokenPrivateKey)
		if err != nil {
			utils.APIErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		refreshToken, err := utils.CreateToken(config.Config.RefreshTokenExpiresIn, user.ID, config.Config.RefreshTokenPrivateKey)
		if err != nil {
			utils.APIErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.SetCookie("access_token", accessToken, config.Config.AccessTokenMaxAge*60, "/", config.Config.Url, false, true)
		ctx.SetCookie("refresh_token", refreshToken, config.Config.RefreshTokenMaxAge*60, "/", config.Config.Url, false, true)
		ctx.SetCookie("logged_in", "true", config.Config.AccessTokenMaxAge*60, "/", config.Config.Url, false, false)

		accessTokenresp := &modelsResponse.AccessTokeResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    time.Now().UTC().Add(config.Config.RefreshTokenExpiresIn).Unix(),
		}
		utils.ObjectToJson(user, &accessTokenresp.User)
		utils.APIResponse(ctx, "Login successfully", http.StatusOK, accessTokenresp)
		return

	default:

		utils.APIResponse(ctx, "Unknown error occured", http.StatusInternalServerError, nil)
	}
}

// Register new user account and return [authCtrl.RegisterResponse]
func (h *Handler) RegisterHandler(ctx *gin.Context) {
	var input modelsInput.RegisterInput

	if ctx.ShouldBind(&input) != nil {
		ctx.JSON(400, gin.H{"error": "Invalid form"})
		return
	}

	errorResponse, errCount := utils.GoValidator(&input)
	if errCount > 0 {
		utils.APIErrorResponse(ctx, http.StatusForbidden, errorResponse)
		return
	}

	_, errorCode := h.service.Register(&input)

	switch errorCode {
	case http.StatusCreated:
		utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, nil)
		return

	case http.StatusConflict:
		utils.APIResponse(ctx, "Email already taken", http.StatusConflict, nil)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(ctx, "Unable to create an account", http.StatusExpectationFailed, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, nil)
	}
}

// Logout the user
func (h *Handler) LogoutHandler(ctx *gin.Context) {
	h.service.Logout(ctx)
	utils.APIResponse(ctx, "Logout successfully", http.StatusOK, nil)
}

// RefreshTokenHandler refresh the token
func (h *Handler) RefreshTokenHandler(ctx *gin.Context) {
	message := "Culd not refresh token"

	cookie, err := ctx.Cookie("refresh_token")
	fmt.Println(cookie)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusForbidden, message)
		return
	}

	sub, err := utils.ValidateToken(cookie, config.Config.RefreshTokenPublicKey)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusForbidden, message)
		return
	}

	user, status := h.service.GetUserByUUID(uuid.MustParse(sub.(string)))
	if status != http.StatusOK {
		utils.APIErrorResponse(ctx, http.StatusForbidden, "User not found")
		return
	}

	accessToken, err := utils.CreateToken(config.Config.AccessTokenExpiresIn, user.ID, config.Config.AccessTokenPrivateKey)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}

	refreshToken, err := utils.CreateToken(config.Config.RefreshTokenExpiresIn, user.ID, config.Config.RefreshTokenPrivateKey)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.SetCookie("access_token", accessToken, config.Config.AccessTokenMaxAge*60, "/", config.Config.Url, false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.Config.RefreshTokenMaxAge*60, "/", config.Config.Url, false, true)
	ctx.SetCookie("logged_in", "true", config.Config.AccessTokenMaxAge*60, "/", config.Config.Url, false, false)

	accessTokenresp := &modelsResponse.AccessTokeResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    time.Now().UTC().Add(config.Config.RefreshTokenExpiresIn).Unix(),
	}
	utils.ObjectToJson(user, &accessTokenresp.User)

	utils.APIResponse(ctx, "Refresh token successfully", http.StatusOK, accessTokenresp)

}
