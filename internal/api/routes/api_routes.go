package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAPIRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/api")
	v1.GET("/ping", Test)
	AddAdminRoute(v1, db)

	AddAuthRoute(v1, db)
	AddUserRoute(v1, db)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
