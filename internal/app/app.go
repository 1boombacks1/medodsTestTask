package app

import (
	"context"
	"d0c/TestTaskBackDev/config"
	"d0c/TestTaskBackDev/helper"
	"d0c/TestTaskBackDev/internal/controllers"
	"d0c/TestTaskBackDev/internal/repo/mongodb"
	"d0c/TestTaskBackDev/internal/services"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// @title Authentication Service
// @version 1.0
// @description This is a service for generating authentication tokens and their refresh

// @contact.name Nikolaev Yakov
// @contact.email nikolaevforbuss@gmail.com

// @host localhost:3003
// @BasePath /
func Run() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config err: %s", err)
	}

	// Database
	log.Info("Initializing mongo...")
	db, err := config.NewMongoDB(cfg.Mongo.Uri, cfg.Mongo.Name)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Main - config.NewMongoDB: %w", err))
	}

	// Repositories
	log.Info("Initializing repositories...")
	sessionRepository := mongodb.NewSessionRepository(db)

	// Services
	log.Info("Initializing services...")
	servicesDeps := services.ServicesDependencies{
		SessionRepo: sessionRepository,
		Hasher:      helper.NewBcryptHasher(cfg.Cost),
		Config:      *cfg,
	}
	services := services.NewServices(servicesDeps)

	// Fiber router
	log.Info("Initializing handlers and routes...")
	app := fiber.New()
	controllers.NewRouter(app, services)

	serverNotify := make(chan error)
	go func() {
		// Fiber server
		log.Info("Starting fiber server...")
		serverNotify <- app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
		close(serverNotify)
	}()

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Info("app - Main - signal: " + s.String())
	case err = <-serverNotify:
		log.Error(fmt.Errorf("app - Main - server notify: %w", err))
	}

	// Graceful shutdown
	log.Info("Gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = app.ShutdownWithContext(ctx); err != nil {
		log.Error(fmt.Errorf("app - Main - app.ShutdownWithContext: %w", err))
	}
}
