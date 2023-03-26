package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Login    string
	Password string
	DBName   string
	Host     string
	Port     string
}

func NewMySqlDB(cfg Config) (*sql.DB, error) {
	req := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Login, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	conn, err := sql.Open("mysql", req)
	if err != nil {
		return nil, err
	}

	//Check for connection
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
