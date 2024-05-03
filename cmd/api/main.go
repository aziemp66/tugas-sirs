package main

import (
	ecg_controller "tugas-sirs/internal/modules/ecg/controller"
	ecg_repository_postgres "tugas-sirs/internal/modules/ecg/repository/postgres"
	ecg_router "tugas-sirs/internal/modules/ecg/router"
	ecg_usecase "tugas-sirs/internal/modules/ecg/usecase"
	"tugas-sirs/internal/pkg/configs"
	pkg_db "tugas-sirs/pkg/db"
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

	pgDb := pkg_db.NewPostgresDB(pkg_db.PostgresDsn{
		Host:     env.PostgresHost,
		Port:     env.PostgresPort,
		User:     env.PostgresUser,
		Password: env.PostgresPassword,
		Db:       env.PostgresDb,
	})

	ecgRepository := ecg_repository_postgres.NewEcgRepositoryPostgres(pgDb)
	ecgUsecase := ecg_usecase.NewEcgUsecase(ecgRepository)
	ecgController := ecg_controller.NewEcgController(ecgUsecase)
	ecg_router.BindEcgRouter(srv.Group("/ecg"), ecgController)

	if err := srv.Run(":" + env.Port); err != nil {
		panic(err)
	}
}
