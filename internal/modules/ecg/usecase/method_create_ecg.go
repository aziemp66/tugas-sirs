package ecg_usecase

import (
	"context"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
)

// Implements EcgUsecase
// TODO: Comment Here
func (ecgUsecaseImplementation *ecgUsecaseImplementation) CreateEcg(ctx context.Context, name string, age int, sex bool, heartFailure int) (id string, err error) {
	heartFailureClass, err := ecg_entity.IntToHeartFailureClass(heartFailure)
	if err != nil {
		return "",err
	}

	id,err = ecgUsecaseImplementation.ecgRepository.CreateEcg(ctx,name,age,sex,heartFailureClass)
	if err != nil {
		return "",err
	}

	return id,nil
}
