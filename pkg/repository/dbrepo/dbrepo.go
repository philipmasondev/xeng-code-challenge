package dbrepo

import (
	"database/sql"
	"web-application-template/pkg/config"
	"web-application-template/pkg/repository"
)

//
type postgresDNRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresDBRepo returns reg to db.
func NewPostgresDBRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDNRepo{
		App: a,
		DB:  conn,
	}
}
