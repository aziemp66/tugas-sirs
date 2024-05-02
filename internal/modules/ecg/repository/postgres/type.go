package ecg_repository_postgres

import "github.com/jmoiron/sqlx"

type ecgRepositoryPostgres struct {
	db *sqlx.DB
}
