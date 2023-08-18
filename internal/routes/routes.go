package routes

import (
	"pomo/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupApiRoutes(router *gin.Engine, handler *handlers.Handler) {
	v1 := router.Group("/api")

	AddAdminRoute(v1, handler)
	AddAuthRoute(v1, handler)
	AddUserRoute(v1, handler)
}
