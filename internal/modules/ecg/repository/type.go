package ecg_repository

import (
	"context"
	"database/sql"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
)

type EcgRepository interface {
	CreateEcg(ctx context.Context, tx *sql.Tx, name string, age int, sex bool, heartFailure ecg_entity.HeartFailureClass) (id string, err error)
	FindEcgByID(ctx context.Context, tx *sql.Tx, id string) (*ecg_entity.Ecg, error)
	FindAllEcg(ctx context.Context, tx *sql.Tx) (*[]*ecg_entity.Ecg, error)
	UpdateEcgByID(ctx context.Context, tx *sql.Tx, id string, name string, age int, sex bool, heartFailure ecg_entity.HeartFailureClass) error
	DeleteEcgByID(ctx context.Context, tx *sql.Tx, id string) error
}
