package api

import (
	apiHandlers "pomo/internal/api/handlers"
	"pomo/internal/api/middlewares"

	"github.com/gin-gonic/gin"
)

func AddAuthRoute(rg *gin.RouterGroup, handler *apiHandlers.Handler) {

	router := rg.Group("/auth")
	router.POST("/login", handler.LoginHandler)
	router.POST("/register", handler.RegisterHandler)
	router.GET("/verify-email/:code", handler.VerifyEmailHandler)
	router.POST("/resend-verification-email", handler.ResendEmailVerificationHandler)
	router.POST("/forgot-password", handler.ForgotPasswordHandler)
	router.PATCH("/reset-password", handler.ResetPasswordHandler)

	router.Use(middlewares.Auth())
	{
		router.POST("/logout", handler.LogoutHandler)
		router.POST("/refresh", handler.RefreshTokenHandler)
		router.POST("/delete-account", handler.DeleteAccountHandler)
	}
}
