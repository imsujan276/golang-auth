package routes

import (
	"pomo/internal/handlers"
	"pomo/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AddAdminRoute(rg *gin.RouterGroup, handler *handlers.Handler) {

	router := rg.Group("/admin", middlewares.Auth(), middlewares.RoleMiddleware("admin"))
	{
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

}
