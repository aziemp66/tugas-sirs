package pkg_db

import (
	"fmt"
	pkg_logger "tugas-sirs/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(pgDsn PostgresDsn) *sqlx.DB {
	db, err := sqlx.Open("postgres", pgDsn.ToString())
	if err != nil {
		pkg_logger.Fatal(err.Error())
	}

	return db
}

type PostgresDsn struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

func (p PostgresDsn) ToString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", p.Host, p.User, p.Password, p.Db, p.Port)
}
