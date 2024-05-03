package ecg_usecase

import (
	"context"
	ecg_response "tugas-sirs/internal/modules/ecg/model/response"
)

// Implements EcgUsecase
// TODO: Comment Here
func (ecgUsecaseImplementation *ecgUsecaseImplementation) FindAllEcg(ctx context.Context) (*[]*ecg_response.EcgResponse, error) {		
	ecgs,err := ecgUsecaseImplementation.ecgRepository.FindAllEcgs(ctx)
	if err != nil {
		return nil, err
	}

	ecgsRes := new([]*ecg_response.EcgResponse)
	
	for _, v := range *ecgs {
		res := &ecg_response.EcgResponse{
			ID: v.ID,
			Name: v.Name,
			Age: v.Age,
			Sex: v.Sex,
			HeartFailure: string(v.HeartFailure),
		}

		*ecgsRes = append(*ecgsRes, res)
	}

	return ecgsRes,nil
}
