package ecg_controller

import (
	"net/http"
	ecg_request "tugas-sirs/internal/modules/ecg/model/request"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"

	"github.com/gin-gonic/gin"
)

// Implements EcgController
// TODO: Comment Here
func (ecgControllerImplementation *ecgControllerImplementation) CreateEcg(ctx *gin.Context) {
	req := new(ecg_request.CreateEcgRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	id, err := ecgControllerImplementation.ecgUsecase.CreateEcg(ctx, req.Name, req.Age, req.Sex, req.HeartFailure)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, pkg_http_wrapper.NewResponseWithValue(http.StatusCreated, "Successfully Created Ecg Record", id))
}
