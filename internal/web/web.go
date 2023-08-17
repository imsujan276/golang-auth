package web

import (
	webHandlers "pomo/internal/web/handlers"
	web "pomo/internal/web/routes"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns the web router.
func SetupRouter(router *gin.Engine, handler *webHandlers.Handler) *gin.Engine {
	// Attach web routes defined in web/routes/web_routes.go
	web.AddWebRoutes(router, handler)

	return router
}
