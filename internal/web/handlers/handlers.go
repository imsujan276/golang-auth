package webHandlers

import (
	"pomo/internal/config"
	"pomo/internal/models"
	"pomo/internal/services"
	"pomo/internal/utils"
	"pomo/templates"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   services.Service
	appConfig *config.AppConfig
}

func NewHandler(service services.Service, appConfig *config.AppConfig) *Handler {
	return &Handler{
		service:   service,
		appConfig: appConfig,
	}
}

func (h *Handler) Home(ctx *gin.Context) {
	// render.Template(w, r, "home.page.html", &models.TemplateData{})
	utils.HTMLResponse(ctx, templates.Home, &models.TemplateData{})
}
