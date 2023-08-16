package api

import (
	usercontroller "pomo/internal/api/controllers/user-controller"
	userhandler "pomo/internal/api/handlers/user-handler"
	"pomo/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUserRoute(rg *gin.RouterGroup, db *gorm.DB) {
	userRepo := usercontroller.NewUserRepository(db)
	userService := usercontroller.NewUserService(userRepo)
	userHandler := userhandler.NewUserHandler(userService)

	router := rg.Group("/user")
	router.GET("/:uuid", userHandler.GetUserByUUID)

	router.Use(middlewares.Auth(), middlewares.RoleMiddleware("user"))
	{
		router.GET("/me", userHandler.GetMe)
	}
}
