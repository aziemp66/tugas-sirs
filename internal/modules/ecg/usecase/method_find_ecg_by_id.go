package ecg_usecase

import (
	"context"
	ecg_response "tugas-sirs/internal/modules/ecg/model/response"
)

// FindEcgByID retrieves an ECG record by its ID and converts it to an EcgResponse object.
// It takes the context and the ID of the ECG record as input.
// If the ECG record is not found, it returns nil and a NotFound error.
// If an error occurs during the retrieval process, it returns nil and the error.
// Otherwise, it returns the EcgResponse object and nil.
func (ecgUsecaseImplementation *ecgUsecaseImplementation) FindEcgByID(ctx context.Context, id string) (*ecg_response.EcgResponse, error) {		
	ecg, err := ecgUsecaseImplementation.ecgRepository.FindEcgByID(ctx, id)
	if err != nil {
		return nil, err
	}

	ecgRes := &ecg_response.EcgResponse{
		ID:           ecg.ID,
		Name:         ecg.Name,
		Age:          ecg.Age,
		Sex:          ecg.Sex,
		HeartFailure: string(ecg.HeartFailure),
	}

	return ecgRes, nil
}
