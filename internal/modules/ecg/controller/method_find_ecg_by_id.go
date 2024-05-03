package ecg_controller

import (
	"net/http"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// Implements EcgController
// TODO: Comment Here
func (ecgControllerImplementation *ecgControllerImplementation) FindEcgByID(ctx *gin.Context) {
	id := ctx.Param("id")

	ecg, err := ecgControllerImplementation.ecgUsecase.FindEcgByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, pkg_http_wrapper.NewResponseWithValue(http.StatusOK, "Successfully Get Ecg Record By ID", ecg))
}
