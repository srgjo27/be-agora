package main

import (
	"log"

	"github.com/srgjo27/agora/internal/config"
	"github.com/srgjo27/agora/internal/handler/http"
	"github.com/srgjo27/agora/internal/repository/postgres"
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

	userUsecase := usecase.NewUserUsecase(userRepo)

	userHandler := http.NewUserHandler(userUsecase)

	router := http.NewRouter(userHandler)

	serverAddress := ":" + cfg.APIPort
	log.Printf("Menjalankan server di %s", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
