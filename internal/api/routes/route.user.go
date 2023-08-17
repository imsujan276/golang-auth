package api

import (
	apiHandlers "pomo/internal/api/handlers"
	"pomo/internal/api/middlewares"

	"github.com/gin-gonic/gin"
)

func AddUserRoute(rg *gin.RouterGroup, handler *apiHandlers.Handler) {

	router := rg.Group("/user")
	router.GET("/:uuid", handler.GetUserByUUID)

	router.Use(middlewares.Auth(), middlewares.RoleMiddleware("user"))
	{
		router.GET("/me", handler.GetMe)
	}
}
