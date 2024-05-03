package ecg_controller

import (
	"net/http"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// DeleteEcgByID handles the deletion of an ECG record by its ID.
// It takes the Gin context as input.
// It extracts the ECG ID from the context parameters.
// If an error occurs during the deletion process, it returns the error.
// Otherwise, it returns a JSON response with the status code 200 and a success message.
func (ecgControllerImplementation *ecgControllerImplementation) DeleteEcgByID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ecgControllerImplementation.ecgUsecase.DeleteEcgByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, pkg_http_wrapper.NewResponse(http.StatusOK, "Successfully Delete Ecg Record"))
}
