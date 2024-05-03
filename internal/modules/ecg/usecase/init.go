package ecg_usecase

import ecg_repository "tugas-sirs/internal/modules/ecg/repository"

func NewEcgUsecase(ecgRepository ecg_repository.EcgRepository) EcgUsecase {
	return &ecgUsecaseImplementation{
		ecgRepository: ecgRepository,
	}
}