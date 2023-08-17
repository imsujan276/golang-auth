package api

import (
	apiHandlers "pomo/internal/api/handlers"
	api "pomo/internal/api/routes"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns the API router.
func SetupRouter(router *gin.Engine, handler *apiHandlers.Handler) *gin.Engine {
	api.AddAPIRoutes(router, handler)
	return router
}
