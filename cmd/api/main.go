package main

import (
	pkg_http "tugas-sirs/pkg/http"
	pkg_http_middleware "tugas-sirs/pkg/http/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	srv := pkg_http.NewHTTPServer(gin.DebugMode)

	srv.Use(
		pkg_http_middleware.TraceIdAssignmentMiddleware(),
		pkg_http_middleware.LogHandler(pkg_http.LOGFILE),
		gin.Recovery(),
		gin.Logger(),
		pkg_http_middleware.ErrorHandler(),
	)

}