package ecg_router

import (
	ecg_controller "tugas-sirs/internal/modules/ecg/controller"

	"github.com/gin-gonic/gin"
)

func BindEcgRouter(router *gin.RouterGroup,ecgController ecg_controller.EcgController)  {
	router.POST("",ecgController.CreateEcg)
	router.GET("",ecgController.FindAllEcg)
	router.GET("/:id",ecgController.FindEcgByID)
	router.PUT("/:id",ecgController.UpdateEcgByID)
	router.DELETE("/:id",ecgController.DeleteEcg)
}