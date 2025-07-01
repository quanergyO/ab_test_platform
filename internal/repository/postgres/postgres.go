package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	slog.Info(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		slog.Error("error open connection with db", slog.Any("error", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		slog.Error("error ping to db", slog.Any("error", err))
		return nil, err
	}

	return db, nil
}
