package ecg_usecase

import (
	"context"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
)

// CreateEcg creates a new ECG record with the provided parameters.
// It takes the context, name, age, sex, and heart failure classification as input.
// If the conversion from the heart failure integer value to the HeartFailureClass fails, it returns an empty string and the conversion error.
// If an error occurs while creating the ECG record in the repository, it returns an empty string and the error.
// Otherwise, it returns the ID of the newly created ECG record and nil.
func (ecgUsecaseImplementation *ecgUsecaseImplementation) CreateEcg(ctx context.Context, name string, age int, sex bool, heartFailure int) (id string, err error) {
	heartFailureClass, err := ecg_entity.IntToHeartFailureClass(heartFailure)
	if err != nil {
		return "", err
	}

	id, err = ecgUsecaseImplementation.ecgRepository.CreateEcg(ctx, name, age, sex, heartFailureClass)
	if err != nil {
		return "", err
	}

	return id, nil
}
