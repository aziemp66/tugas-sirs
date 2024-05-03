package ecg_usecase

import (
	"context"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
)

// UpdateEcgByID updates an ECG record by its ID with the provided information.
// It takes the context, ID of the ECG record, name, age, sex, and heart failure class as input.
// If an error occurs during the update process, it returns the error.
// Otherwise, it returns nil.
func (ecgUsecaseImplementation *ecgUsecaseImplementation) UpdateEcgByID(ctx context.Context, id string, name string, age int, sex bool, heartFailure int) error {		
	heartFailureClass, err := ecg_entity.IntToHeartFailureClass(heartFailure)
	if err != nil {
		return err
	}

	err = ecgUsecaseImplementation.ecgRepository.UpdateEcgByID(ctx, id, name, age, sex, heartFailureClass)
	if err != nil {
		return err
	}

	return nil
}
