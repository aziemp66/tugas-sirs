package ecg_usecase

import (
	"context"
	ecg_response "tugas-sirs/internal/modules/ecg/model/response"
)

// Implements EcgUsecase
// TODO: Comment Here
func (ecgUsecaseImplementation *ecgUsecaseImplementation) FindEcgByID(ctx context.Context, id string) (*ecg_response.EcgResponse, error) {		
	ecg,err := ecgUsecaseImplementation.ecgRepository.FindEcgByID(ctx,id)
	if err != nil {
		return nil, err
	}

	ecgRes := &ecg_response.EcgResponse{
		ID: ecg.ID,
		Name: ecg.Name,
		Age: ecg.Age,
		Sex: ecg.Sex,
		HeartFailure: string(ecg.HeartFailure),
	}

	return ecgRes,nil
}
