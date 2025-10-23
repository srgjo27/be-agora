package main

import (
	"log"

	"github.com/srgjo27/agora/internal/config"
	"github.com/srgjo27/agora/internal/handler/http"
	"github.com/srgjo27/agora/internal/repository/postgres"
	"github.com/srgjo27/agora/internal/service"
	"github.com/srgjo27/agora/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Tidak bisa memuat config: %v", err)
	}

	db := postgres.ConnectDB(&cfg)
	log.Printf("Berhasil terhubung ke DB: %s di host %s", cfg.DBName, cfg.DBHost)

	userRepo := postgres.NewPostgresUserRepo(db)

	tokenSvc := service.NewTokenService(&cfg)

	userUsecase := usecase.NewUserUsecase(userRepo, tokenSvc)

	userHandler := http.NewUserHandler(userUsecase, &cfg)

	router := http.NewRouter(userHandler)

	serverAddress := ":" + cfg.APIPort
	log.Printf("Menjalankan server di %s", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
