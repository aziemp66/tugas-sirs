package ecg_usecase

import (
	"context"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
)

// Implements EcgUsecase
// TODO: Comment Here
func (ecgUsecaseImplementation *ecgUsecaseImplementation) UpdateEcgByID(ctx context.Context, id string, name string, age int, sex bool, heartFailure int) error {		
	heartFailureClass, err := ecg_entity.IntToHeartFailureClass(heartFailure)
	if err != nil {
		return err
	}

	err = ecgUsecaseImplementation.ecgRepository.UpdateEcgByID(ctx,id,name,age,sex,heartFailureClass)
	if err != nil {
		return err
	}

	return nil
}
