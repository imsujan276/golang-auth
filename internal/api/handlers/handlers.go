package apiHandlers

import (
	"pomo/internal/config"
	"pomo/internal/services"
)

type Handler struct {
	appConfig *config.AppConfig
	service   services.Service
}

func NewHandler(service services.Service, appConfig *config.AppConfig) *Handler {
	return &Handler{
		service:   service,
		appConfig: appConfig,
	}
}
