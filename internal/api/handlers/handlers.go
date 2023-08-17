package apiHandlers

import "pomo/internal/services"

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{
		service: service,
	}
}
