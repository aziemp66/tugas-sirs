package ecg_controller

import (
	"net/http"
	ecg_request "tugas-sirs/internal/modules/ecg/model/request"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// UpdateEcgByID handles the update of an ECG record by its ID.
// It takes the Gin context as input.
// It extracts the ECG ID and update request data from the request parameters.
// It binds the request data to the UpdateEcgByID struct.
// It delegates the update process to the EcgUsecase.
// If an error occurs during the update process, it returns the error.
// Otherwise, it returns a JSON response with the status code 200 indicating successful update.
func (ecgControllerImplementation *ecgControllerImplementation) UpdateEcgByID(ctx *gin.Context) {
	id := ctx.Param("id")
	req := new(ecg_request.UpdateEcgByID)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	err := ecgControllerImplementation.ecgUsecase.UpdateEcgByID(ctx, id, req.Name, req.Age, req.Sex, req.HeartFailure)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, pkg_http_wrapper.NewResponse(http.StatusOK, "Successfully Updated Ecg Record by ID"))
}
