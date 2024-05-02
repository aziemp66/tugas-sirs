package ecg_controller

import "github.com/gin-gonic/gin"

type EcgController interface {
	CreateEcg(ctx *gin.Context)
	FindEcgByID(ctx *gin.Context)
	FindAllEcg(ctx *gin.Context)
	UpdateEcgByID(ctx *gin.Context)
	DeleteEcg(ctx *gin.Context)
}