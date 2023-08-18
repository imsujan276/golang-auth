package routes

import (
	"pomo/internal/handlers"
	"pomo/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AddUserRoute(rg *gin.RouterGroup, handler *handlers.Handler) {

	router := rg.Group("/user")
	router.GET("/:uuid", handler.GetUserByUUID)

	router.Use(middlewares.Auth(), middlewares.RoleMiddleware("user"))
	{
		router.GET("/me", handler.GetMe)
	}
}
