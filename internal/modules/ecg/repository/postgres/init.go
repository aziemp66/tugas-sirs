package ecg_repository_postgres

import (
	"github.com/jmoiron/sqlx"
)

func NewEcgRepositoryPostgres(db *sqlx.DB) *ecgRepositoryPostgres {
	return &ecgRepositoryPostgres{
		db: db,
	}
}
