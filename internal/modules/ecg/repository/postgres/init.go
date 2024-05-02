package ecg_repository_postgres

import (
	ecg_repository "tugas-sirs/internal/modules/ecg/repository"

	"github.com/jmoiron/sqlx"
)

func NewEcgRepositoryPostgres(db *sqlx.DB) ecg_repository.EcgRepository {
	return &ecgRepositoryPostgres{
		db: db,
	}
}
