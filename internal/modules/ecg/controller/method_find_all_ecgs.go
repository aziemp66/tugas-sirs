package ecg_controller

import (
	"net/http"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// FindAllEcgs handles the retrieval of all ECG records.
// It takes the Gin context as input.
// It delegates the retrieval process to the EcgUsecase.
// If an error occurs during the retrieval process, it returns the error.
// Otherwise, it returns a JSON response with the status code 200 and the retrieved ECG records.
func (ecgControllerImplementation *ecgControllerImplementation) FindAllEcgs(ctx *gin.Context) {
	ecgs, err := ecgControllerImplementation.ecgUsecase.FindAllEcgs(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, pkg_http_wrapper.NewResponseWithValue(http.StatusOK, "Successfully Get All Ecgs Record", ecgs))
}
