package web

import (
	web "pomo/internal/web/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter initializes and returns the web router.
func SetupRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	// Attach web routes defined in web/routes/web_routes.go
	web.AddWebRoutes(router, db)

	return router
}
