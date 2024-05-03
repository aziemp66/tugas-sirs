package ecg_controller

import ecg_usecase "tugas-sirs/internal/modules/ecg/usecase"

func NewEcgController(ecgUsecase ecg_usecase.EcgUsecase) EcgController {
	return &ecgControllerImplementation{
		ecgUsecase: ecgUsecase,
	}
}
