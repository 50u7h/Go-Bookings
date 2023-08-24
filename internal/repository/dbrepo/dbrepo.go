package dbrepo

import (
	"database/sql"
	"github.com/50u7h/Go-Bookings/internal/config"
	"github.com/50u7h/Go-Bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
