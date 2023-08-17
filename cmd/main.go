package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pomo/internal/api"
	apiHandlers "pomo/internal/api/handlers"
	"pomo/internal/config"
	"pomo/internal/database"
	"pomo/internal/render"
	repo "pomo/internal/repository"
	"pomo/internal/services"
	"pomo/internal/web"
	webHandlers "pomo/internal/web/handlers"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setup()
	serve(router, config.Config)
}

// setup sets up the server
func setup() *gin.Engine {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
		os.Exit(1)
	}

	config.Config = appConfig
	router := gin.Default()

	database.Connection(config.Config)
	render.NewRenderer(config.Config)

	setupRoutes(router)
	setupStaticFiles(router)

	if config.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return router
}

// setupRoutes sets up the routes for web and api
func setupRoutes(router *gin.Engine) {
	repository := repo.NewRepository(database.DB)
	service := services.NewService(repository)
	apiHandler := apiHandlers.NewHandler(service)
	webHandler := webHandlers.NewHandler(service)

	api.SetupRouter(router, apiHandler)
	web.SetupRouter(router, webHandler)

}

// setupStaticFiles sets up the static files
func setupStaticFiles(router *gin.Engine) {
	router.Static("/static/images", "./static/images")
}

// serve start server using Graceful Shutdown
func serve(router *gin.Engine, appConfig *config.AppConfig) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.ServerPort),
		Handler: router,
	}

	go func() {
		log.Printf("Server listening on :%d", appConfig.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Trap SIGINT and SIGTERM signals to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a context for graceful shutdown with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
