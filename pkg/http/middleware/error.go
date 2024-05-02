package pkg_http_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	errorPkg "tugas-sirs/pkg/error"
	pkg_http_wrapper "tugas-sirs/pkg/http/wrapper"
	pkg_logger "tugas-sirs/pkg/logger"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0]
		// if err can be casted to ClientError, then it is a client error
		if clientError, ok := err.Err.(*errorPkg.ClientError); ok {
			pkg_logger.Error(clientError.Raw.Error(), zap.Int("Code", clientError.Code))
			c.JSON(clientError.Code, pkg_http_wrapper.NewError(clientError.Code,clientError.Message))
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			pkg_logger.Error(err.Error(), zap.Any("Error", err))
			c.JSON(400, pkg_http_wrapper.NewError(http.StatusBadRequest,err.Err.Error()))
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			pkg_logger.Error(err.Error(), zap.Any("Error", err))
			c.JSON(500, pkg_http_wrapper.NewError(http.StatusInternalServerError,"Internal Server Error"))
			return
		}

		pkg_logger.Error(err.Error(), zap.Any("Error", err))
		c.JSON(500, pkg_http_wrapper.NewError(http.StatusInternalServerError,"Internal Server Error"))
	}
}
