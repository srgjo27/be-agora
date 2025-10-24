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

	log.Printf("[CONFIG] Loaded: APIPort=%s, AccessTokenDuration=%d, SecretKeyIsSet=%t", cfg.APIPort, cfg.AccessTokenDurationMinutes, cfg.JWTSecretKey != "")

	db := postgres.ConnectDB(&cfg)
	log.Printf("Berhasil terhubung ke DB: %s di host %s", cfg.DBName, cfg.DBHost)

	userRepo := postgres.NewPostgresUserRepo(db)
	categoryRepo := postgres.NewPostgresCategoryRepo(db)

	tokenSvc := service.NewTokenService(&cfg)

	userUsecase := usecase.NewUserUsecase(userRepo, tokenSvc)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	userHandler := http.NewUserHandler(userUsecase, &cfg)
	categoryHandler := http.NewCategoryHandler(categoryUsecase)

	authMiddleware := http.NewAuthMiddleware(tokenSvc)

	router := http.NewRouter(
		userHandler,
		authMiddleware,
		categoryHandler,
	)

	serverAddress := ":" + cfg.APIPort
	log.Printf("Menjalankan server di %s", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
