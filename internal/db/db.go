package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "host=localhost user=jorgeemiliano password=Jorge41304254 dbname=turismo_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida con éxito")
	return nil
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("Error al obtener la conexión SQL: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error al cerrar la conexión a la base de datos: %v", err)
		} else {
			log.Println("Conexión a la base de datos cerrada con éxito")
		}
	}
}
