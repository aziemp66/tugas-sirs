package ecg_usecase

import "context"

// Implements EcgUsecase
// TODO: Comment Here
func (ecgUsecaseImplementation *ecgUsecaseImplementation) DeleteEcgByID(ctx context.Context, id string) error {		
	err := ecgUsecaseImplementation.ecgRepository.DeleteEcgByID(ctx,id)
	if err != nil {
		return err
	}

	return nil
}