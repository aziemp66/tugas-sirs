package ecg_usecase

import (
	"context"
	ecg_response "tugas-sirs/internal/modules/ecg/model/response"
	ecg_repository "tugas-sirs/internal/modules/ecg/repository"
)

type ecgUsecaseImplementation struct {
	ecgRepository ecg_repository.EcgRepository
}

type EcgUsecase interface {
	CreateEcg(ctx context.Context, name string, age int, sex bool, heartFailure int) (id string, err error)
	FindEcgByID(ctx context.Context, id string) (*ecg_response.EcgResponse, error)
	FindAllEcgs(ctx context.Context) (*[]*ecg_response.EcgResponse, error)
	UpdateEcgByID(ctx context.Context, id string, name string, age int, sex bool, heartFailure int) error
	DeleteEcgByID(ctx context.Context, id string) error
}
