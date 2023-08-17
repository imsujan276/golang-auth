package api

import (
	apiHandlers "pomo/internal/api/handlers"
	"pomo/internal/api/middlewares"

	"github.com/gin-gonic/gin"
)

func AddAdminRoute(rg *gin.RouterGroup, handler *apiHandlers.Handler) {

	router := rg.Group("/admin", middlewares.Auth(), middlewares.RoleMiddleware("admin"))
	{
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

}
