package ecg_repository

import (
	"context"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
)

type EcgRepository interface {
	CreateEcg(ctx context.Context, name string, age int, sex bool, heartFailure ecg_entity.HeartFailureClass) (id string, err error)
	FindEcgByID(ctx context.Context, id string) (*ecg_entity.Ecg, error)
	FindAllEcg(ctx context.Context) ([]*ecg_entity.Ecg, error)
	UpdateEcgByID(ctx context.Context, id string, name string, age int, sex bool, heartFailure ecg_entity.HeartFailureClass) error
	DeleteEcgByID(ctx context.Context, id string) error
}
