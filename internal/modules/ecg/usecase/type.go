package ecg_usecase

import (
	"context"
	ecg_response "tugas-sirs/internal/modules/ecg/model/response"
)

type EcgUsecase interface {
	CreateEcg(ctx context.Context, name string, age int, sex bool, heartFailure int) error
	FindEcgByID(ctx context.Context, id string) (*ecg_response.EcgResponse, error)
	FindAllEcg(ctx context.Context) ([]*ecg_response.EcgResponse, error)
	UpdateEcgByID(ctx context.Context, id string, name string, age int, sex bool, heartFailure int) error
	DeleteEcgByID(ctx context.Context, id string) error
}
