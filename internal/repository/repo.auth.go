package repo

import (
	"net/http"
	"pomo/internal/config"
	"pomo/internal/models"
	"pomo/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// The LoginRepository method of the repository struct implements the LoginRepository method
// from the Repository interface.
func (r *repository) Login(input *models.UserModel) (*models.UserModel, int) {
	var user models.UserModel
	db := r.db.Model(&user)
	checkAccount := db.Select("*").Where("email=?", strings.ToLower(input.Email)).Find(&user)

	if checkAccount.RowsAffected == 0 {
		return nil, http.StatusNotFound
	}

	if !user.Verified {
		return nil, http.StatusForbidden
	}
	// check if the password matches
	verifyPassword := utils.VerifyPassword(user.Password, input.Password)

	if verifyPassword != nil {
		return nil, http.StatusUnauthorized
	}

	return &user, http.StatusAccepted
}

func (r *repository) Register(input *models.UserModel) (*models.UserModel, int) {
	var user models.UserModel

	//  check if user exists
	checkUserAccount := r.db.Select("*").Where("Email=?", input.Email).Find(&user)
	if checkUserAccount.RowsAffected > 0 {
		return nil, http.StatusConflict
	}

	//  if not then create the user into db
	createUser := r.db.Create(&input)

	if createUser.Error != nil {
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}

func (r *repository) Logout(ctx *gin.Context) int {
	ctx.SetCookie("access_token", "", -1, "/", config.Config.Url, false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", config.Config.Url, false, true)
	ctx.SetCookie("logged_in", "", -1, "/", config.Config.Url, false, false)
	ctx.Set("user", nil)
	ctx.Set("role", nil)
	return http.StatusOK
}

func (r *repository) DeleteUser(user *models.UserModel) error {
	return r.db.Delete(&user).Error
}
