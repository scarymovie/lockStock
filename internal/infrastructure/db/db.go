package db

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL драйвер
	"lockStock/pkg/logger"
	"log"
)

func NewDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=lock-stock-dev-db port=5432 user=db_user password=db_password dbname=db_database sslmode=disable")
	if err != nil {
		logger.Logger.Fatalf("error opening database: %s", err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		logger.Logger.Fatalf("database not reachable: %s", err.Error())
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}
