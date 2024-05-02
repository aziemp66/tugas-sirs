package configs

type config struct {
	Port             string `env:"APP_PORT"`
	AppEnv           string `env:"APP_ENV"`
	LogPath          string `env:"LOG_PATH"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
	PostgresDb       string `env:"POSTGRES_DB"`
}
