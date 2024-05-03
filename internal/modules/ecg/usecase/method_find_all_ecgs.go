package ecg_usecase

import (
	"context"
	ecg_response "tugas-sirs/internal/modules/ecg/model/response"
)

// FindAllEcg retrieves all ECG records and converts them to EcgResponse objects.
// It takes the context as input.
// If an error occurs during the retrieval process, it returns nil and the error.
// Otherwise, it returns a slice of EcgResponse pointers and nil.
func (ecgUsecaseImplementation *ecgUsecaseImplementation) FindAllEcgs(ctx context.Context) (*[]*ecg_response.EcgResponse, error) {		
	ecgs, err := ecgUsecaseImplementation.ecgRepository.FindAllEcgs(ctx)
	if err != nil {
		return nil, err
	}

	ecgsRes := new([]*ecg_response.EcgResponse)
	
	for _, v := range *ecgs {
		res := &ecg_response.EcgResponse{
			ID:           v.ID,
			Name:         v.Name,
			Age:          v.Age,
			Sex:          v.Sex,
			HeartFailure: string(v.HeartFailure),
		}

		*ecgsRes = append(*ecgsRes, res)
	}

	return ecgsRes, nil
}
