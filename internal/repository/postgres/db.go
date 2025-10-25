package postgres

import (
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/config"
)

func ConnectDB(cfg *config.Config) *sqlx.DB {
	dsn := cfg.DSN()

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("[ERROR]: Gagal terhubung ke database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("[ERROR]: Gagal ping database: %v", err)
	}

	log.Printf("[SUCCESS]: Koneksi database berhasil!")
	return db
}
