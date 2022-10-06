package db

import (
	"database/sql"
	"fmt"
	"github.com/ValikoDorodnov/go_sample/internal/config"
	_ "github.com/lib/pq"
)

func Init(config config.DbConfig) (*sql.DB, error) {

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.User, config.Pass, config.Name, config.SllMode)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	return db, err
}
