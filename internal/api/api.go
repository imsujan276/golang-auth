package api

import (
	api "pomo/internal/api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter initializes and returns the API router.
func SetupRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {

	// Attach API routes defined in api/routes/api_routes.go
	api.AddAPIRoutes(router, db)

	// Add middleware like authentication
	// router.Use(AuthMiddleware)

	return router
}
