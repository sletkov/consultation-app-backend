package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

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

	sessionStore := sessions.NewCookieStore([]byte(cfg.SessionKey))
	userService := service.NewUserService(userRepository)
	consultationsService := service.NewConsultationService(consultationRepository)

	controller := handlers.NewV1Controller(userService, consultationsService, sessionStore)
	router := controller.InitRoutes()

	if err := http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.Port), router); err != nil {
		log.Fatal("Cannot start server", err)
	}
}
