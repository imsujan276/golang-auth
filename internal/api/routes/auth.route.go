package api

import (
	authCtrl "pomo/internal/api/controllers/auth-controller"
	usercontroller "pomo/internal/api/controllers/user-controller"
	authhandler "pomo/internal/api/handlers/auth-handler"
	"pomo/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAuthRoute(rg *gin.RouterGroup, db *gorm.DB) {
	userRepo := usercontroller.NewUserRepository(db)
	userService := usercontroller.NewUserService(userRepo)

	authRepository := authCtrl.NewAuthRepository(db)
	authService := authCtrl.NewAuthService(authRepository)
	authHandler := authhandler.NewAuthHandler(authService, userService)

	router := rg.Group("/auth")
	router.POST("/login", authHandler.LoginHandler)
	router.POST("/register", authHandler.RegisterHandler)

	router.Use(middlewares.Auth())
	{
		router.POST("/logout", authHandler.LogoutHandler)
		router.POST("/refresh", authHandler.RefreshTokenHandler)
	}
}
