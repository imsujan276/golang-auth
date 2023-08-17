package web

import (
	webHandlers "pomo/internal/web/handlers"

	"github.com/gin-gonic/gin"
)

func AddWebRoutes(router *gin.Engine, handler *webHandlers.Handler) {
	router.GET("/", Test)
	// router.GET("/", handler.Home)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Working..."})
}
