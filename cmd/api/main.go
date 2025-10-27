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
		log.Fatalf("[ERROR]: Tidak bisa memuat config: %v", err)
	}

	log.Printf("[INFO]: APIPort=%s, AccessTokenDuration=%d, SecretKeyIsSet=%t", cfg.APIPort, cfg.AccessTokenDurationMinutes, cfg.JWTSecretKey != "")

	db := postgres.ConnectDB(&cfg)
	log.Printf("[SUCCESS]: Berhasil terhubung ke DB: %s di host %s", cfg.DBName, cfg.DBHost)

	userRepo := postgres.NewPostgresUserRepo(db)
	categoryRepo := postgres.NewPostgresCategoryRepo(db)
	threadRepo := postgres.NewPostgresThreadRepo(db)
	postRepo := postgres.NewPostgresPostRepo(db)
	voteRepo := postgres.NewPostgresVoteRepo(db)

	tokenSvc := service.NewTokenService(&cfg)

	userUsecase := usecase.NewUserUsecase(userRepo, tokenSvc)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	threadUsecase := usecase.NewThreadUsecase(threadRepo, categoryRepo)
	postUsecase := usecase.NewPostUsecase(postRepo, threadRepo)
	voteUsecase := usecase.NewVoteUsecase(db, voteRepo, threadRepo)

	userHandler := http.NewUserHandler(userUsecase, &cfg)
	categoryHandler := http.NewCategoryHandler(categoryUsecase)
	threadHandler := http.NewThreadHandler(threadUsecase)
	postHandler := http.NewPostHandler(postUsecase)
	voteHandler := http.NewVoteHandler(voteUsecase)

	authMiddleware := http.NewAuthMiddleware(tokenSvc)

	router := http.NewRouter(
		userHandler,
		authMiddleware,
		categoryHandler,
		threadHandler,
		postHandler,
		voteHandler,
	)

	serverAddress := ":" + cfg.APIPort
	log.Printf("[SUCCESS]: Menjalankan server di %s", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("[ERROR]: Gagal menjalankan server: %v", err)
	}
}
