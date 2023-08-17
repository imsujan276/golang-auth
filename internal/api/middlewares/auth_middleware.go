package middlewares

import (
	"fmt"
	"net/http"
	"pomo/internal/config"
	"pomo/internal/database"
	"pomo/internal/models"
	"pomo/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware to authenticate API requests.
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			utils.APIErrorResponse(ctx, http.StatusUnauthorized, "You are not logged in")
			return
		}

		sub, err := utils.ValidateToken(access_token, config.Config.AccessTokenPublicKey)
		if err != nil {
			utils.APIErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		var user models.UserModel
		result := database.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			utils.APIErrorResponse(ctx, http.StatusForbidden, "User not found")
			return
		}

		ctx.Set("user", user)
		ctx.Set("role", user.Role)
		ctx.Next()
	}
}
