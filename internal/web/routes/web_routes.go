package web

import (
	webHandlers "pomo/internal/web/handlers"

	"github.com/gin-gonic/gin"
)

func AddWebRoutes(router *gin.Engine, handler *webHandlers.Handler) {
	router.GET("/", handler.Home)
}
