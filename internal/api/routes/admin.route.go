package api

import (
	"pomo/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAdminRoute(rg *gin.RouterGroup, db *gorm.DB) {

	router := rg.Group("/admin", middlewares.Auth(), middlewares.RoleMiddleware("admin"))
	{
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

}
