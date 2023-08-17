package webHandlers

import (
	"net/http"
	"pomo/internal/config"
	"pomo/internal/models"
	"pomo/internal/render"
	"pomo/internal/services"
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

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.html", &models.TemplateData{})
}
