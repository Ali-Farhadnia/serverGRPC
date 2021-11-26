package bookdb

import (
	"database/sql"
)

type BookDB struct {
	Config DBConfig
	Db     *sql.DB
}
type DBConfig struct {
	User     string
	Password string
	Sslmode  string
	Host     string
	Port     string
	DbName   string
}
