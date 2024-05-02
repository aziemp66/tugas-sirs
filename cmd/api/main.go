package main

import (
	"tugas-sirs/internal/pkg/configs"
	pkg_http "tugas-sirs/pkg/http"
	pkg_http_middleware "tugas-sirs/pkg/http/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	env := configs.LoadConfig()

	srv := pkg_http.NewHTTPServer(env.AppEnv)

	srv.Use(
		pkg_http_middleware.TraceIdAssignmentMiddleware(),
		pkg_http_middleware.LogHandler(env.LogPath),
		gin.Recovery(),
		gin.Logger(),
		pkg_http_middleware.ErrorHandler(),
	)

	if err := srv.Run(":" + env.Port); err != nil {
		panic(err)
	}
}
