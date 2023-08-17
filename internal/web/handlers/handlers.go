package webHandlers

import (
	"net/http"
	"pomo/internal/models"
	"pomo/internal/render"
	"pomo/internal/services"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.html", &models.TemplateData{})
}
