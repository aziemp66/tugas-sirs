package ecg_controller

import (
	"net/http"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// Implements EcgController
// TODO: Comment Here
func (ecgControllerImplementation *ecgControllerImplementation) DeleteEcgByID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ecgControllerImplementation.ecgUsecase.DeleteEcgByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, pkg_http_wrapper.NewResponse(http.StatusOK, "Successfully Delete Ecg Record"))
}
