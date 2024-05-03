package ecg_controller

import (
	"net/http"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// Implements EcgController
// TODO: Comment Here
func (ecgControllerImplementation *ecgControllerImplementation) FindAllEcgs(ctx *gin.Context) {
	ecgs, err := ecgControllerImplementation.ecgUsecase.FindAllEcgs(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, pkg_http_wrapper.NewResponseWithValue(http.StatusOK, "Successfully Get All Ecgs Record", ecgs))
}
