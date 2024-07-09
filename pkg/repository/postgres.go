package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	salesListsTable = "sales_lists"
	usersListsTable = "users_lists"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	fmt.Println(cfg.Host)
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.DBName, "postgres", cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
