package ecg_controller

import (
	ecg_usecase "tugas-sirs/internal/modules/ecg/usecase"

	"github.com/gin-gonic/gin"
)

type ecgControllerImplementation struct {
	ecgUsecase ecg_usecase.EcgUsecase
}

type EcgController interface {
	CreateEcg(ctx *gin.Context)
	FindEcgByID(ctx *gin.Context)
	FindAllEcgs(ctx *gin.Context)
	UpdateEcgByID(ctx *gin.Context)
	DeleteEcgByID(ctx *gin.Context)
}
