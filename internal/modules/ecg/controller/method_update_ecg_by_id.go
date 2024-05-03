package ecg_controller

import (
	"net/http"
	ecg_request "tugas-sirs/internal/modules/ecg/model/request"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// Implements EcgController
// TODO: Comment Here
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
