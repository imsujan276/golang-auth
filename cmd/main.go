package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pomo/internal/config"
	"pomo/internal/database"
	"pomo/internal/handlers"
	emailModels "pomo/internal/models/email"
	repo "pomo/internal/repository"
	"pomo/internal/routes"
	"pomo/internal/services"
	"syscall"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router, err := setup()

	if err != nil {
		log.Fatalf("Error setting up server: %v", err)
		os.Exit(1)
	}

	defer close(config.Config.MailChannel)
	listenForMail()
	serve(router, config.Config)
}

// setup sets up the server
func setup() (*gin.Engine, error) {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
		os.Exit(1)
	}

	config.Config = appConfig
	config.Config.MailChannel = make(chan emailModels.MailData)

	session := scs.New()
	session.Lifetime = 24 * time.Hour // session lasts for 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = config.Config.Debug // use ssl; set to true in production
	config.Config.Session = session

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.Config.Url}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	err = database.Connection(config.Config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	setupRoutes(router)
	setupStaticFiles(router)

	if config.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return router, nil
}

// setupRoutes sets up the routes for web and api
func setupRoutes(router *gin.Engine) {
	repository := repo.NewRepository(database.DB)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service, config.Config)
	routes.SetupApiRoutes(router, handler)
}

// setupStaticFiles sets up the static files
func setupStaticFiles(router *gin.Engine) {
	router.Static("/static/images", "./static/images")
	router.Static("/uploads/images", "./uploads/images")
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
