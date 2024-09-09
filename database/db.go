package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Verificar la conexión
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging the database: %w", err)
	}

	return db, nil
}

func InitDB() (*sql.DB, error) {
	var err error
	DB, err = ConnectDB("jorgeemiliano", "Jorge41304254#", "tourist_app")
	if err != nil {
		return nil, err
	}
	return DB, nil
}
