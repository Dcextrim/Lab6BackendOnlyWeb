package db

import (
	"series-tracker/pkg/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establece una conexión con la base de datos PostgreSQL.
// Se realizan hasta 10 intentos en caso de fallo, con un retraso de 5 segundos entre cada intento.
// Retorna una instancia de la base de datos y un error en caso de fallo.
func ConnectDB() (*gorm.DB, error) {
	// Cadena de conexión a la base de datos PostgreSQL.
	dsn := "host=db user=user password=password dbname=seriesdb port=5432 sslmode=disable"

	var db *gorm.DB
	var err error

	// Intentar la conexión hasta 10 veces en caso de error.
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break // Salir del bucle si la conexión es exitosa.
		}
		time.Sleep(5 * time.Second) // Esperar 5 segundos antes de intentar nuevamente.
	}

	// Si no se pudo establecer la conexión, retornar el error.
	if err != nil {
		return nil, err
	}

	// Migrar automáticamente la estructura de datos para la tabla 'series'.
	db.AutoMigrate(&models.Series{})

	return db, nil // Retornar la conexión establecida.
}
