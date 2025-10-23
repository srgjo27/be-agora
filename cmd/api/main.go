package main

import (
	"log"

	"github.com/srgjo27/agora/internal/config"
	"github.com/srgjo27/agora/internal/repository/postgres"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Tidak bisa memuat config: %v", err)
	}

	db := postgres.ConnectDB(&cfg)
	defer db.Close()

	log.Printf("Berhasil terhubung ke DB: %s di host %s", cfg.DBName, cfg.DBHost)
}
