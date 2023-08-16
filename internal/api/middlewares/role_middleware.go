package middlewares

import (
	"net/http"
	"pomo/internal/models"
	"pomo/internal/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if requiredRole != role {
			utils.APIErrorResponse(c, http.StatusForbidden, "Permission denied")
			return
		}
		c.Next()
	}
}

func PermissionMiddleware(requiredPermissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		allowed := false

		for _, permission := range requiredPermissions {
			if hasPermission(role, permission) {
				allowed = true
				break
			}
		}

		if !allowed {
			utils.APIErrorResponse(c, http.StatusForbidden, "Permission denied")
			return
		}

		c.Next()
	}
}

func hasPermission(role string, permission string) bool {
	r, found := models.Roles[role]
	if !found {
		return false
	}

	for _, p := range r.Permissions {
		if p == permission {
			return true
		}
	}

	return false
}
