package dbController

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DBFileName string `yaml:"DBFileName"`
}

type (
	Database interface {
		WellIDs() ([]int, error)
	}
	database struct {
		db *sql.DB
	}
)

func NewDatabase(config Config) (Database, error) {

	d := new(database)

	var err error

	d.db, err = sql.Open("sqlite3", config.DBFileName)
	if err != nil {
		return nil, err
	}
	return d, nil
}
