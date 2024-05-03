package ecg_usecase

import "context"

// DeleteEcgByID deletes an ECG record by its ID.
// It takes the context and the ID of the ECG record to be deleted as input.
// If an error occurs during the deletion process, it returns the error.
// Otherwise, it returns nil.
func (ecgUsecaseImplementation *ecgUsecaseImplementation) DeleteEcgByID(ctx context.Context, id string) error {
	err := ecgUsecaseImplementation.ecgRepository.DeleteEcgByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
