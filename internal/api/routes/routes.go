package api

import (
	apiHandlers "pomo/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func AddAPIRoutes(router *gin.Engine, handler *apiHandlers.Handler) {
	v1 := router.Group("/api")
	v1.GET("/ping", Test)
	AddAdminRoute(v1, handler)

	AddAuthRoute(v1, handler)
	AddUserRoute(v1, handler)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
