package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sletkov/consultation-app-backend/internal/emailnotifier"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/sletkov/consultation-app-backend/db/migrations"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/handlers"
	"github.com/sletkov/consultation-app-backend/internal/config"
	"github.com/sletkov/consultation-app-backend/internal/repository"
	"github.com/sletkov/consultation-app-backend/internal/service"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "./.env", "config path")
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	var cfg config.Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	defer db.Close()

	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	if err := goose.Up(db, "db/migrations"); err != nil {
		slog.Warn(fmt.Errorf("making migrations: %w", err).Error())
	}

	slog.Info("migrations were made successfully")

	userRepository := repository.NewUserRepository(db)
	consultationRepository := repository.NewConsultationRepository(db)

	emailNotifier := emailnotifier.New(&cfg)

	sessionStore := sessions.NewCookieStore([]byte(cfg.SessionKey))
	userService := service.NewUserService(userRepository)
	consultationsService := service.NewConsultationService(consultationRepository, emailNotifier)

	controller := handlers.NewV1Controller(userService, consultationsService, sessionStore)
	router := controller.InitRoutes()

	httpSrv := &http.Server{
		Addr:         net.JoinHostPort(cfg.ServerHost, cfg.ServerPort),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Fatal("Cannot start server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Info("Shutting down server...")
	if err := httpSrv.Shutdown(context.Background()); err != nil {
		slog.Error("error occured on server shutting down", err)
	}

	if err := db.Close(); err != nil {
		slog.Error("error occured on db closing", err)
	}

}
