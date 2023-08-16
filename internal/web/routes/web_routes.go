package web

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddWebRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/", Test)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Working..."})
}
