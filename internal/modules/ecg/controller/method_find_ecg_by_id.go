package ecg_controller

import (
	"net/http"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// FindEcgByID handles the retrieval of an ECG record by its ID.
// It takes the Gin context as input.
// It extracts the ECG ID from the request parameters.
// It delegates the retrieval process to the EcgUsecase.
// If an error occurs during the retrieval process, it returns the error.
// Otherwise, it returns a JSON response with the status code 200 and the retrieved ECG record.
func (ecgControllerImplementation *ecgControllerImplementation) FindEcgByID(ctx *gin.Context) {
	id := ctx.Param("id")

	ecg, err := ecgControllerImplementation.ecgUsecase.FindEcgByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, pkg_http_wrapper.NewResponseWithValue(http.StatusOK, "Successfully Get Ecg Record By ID", ecg))
}
