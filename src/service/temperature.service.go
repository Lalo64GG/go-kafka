package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lalo64/go-kafka/src/models"
)


type TemperatureService struct {
	DB *sql.DB
}

func NewTemperatureService(db *sql.DB) *TemperatureService{
	return &TemperatureService{ DB: db }
}


func (s *TemperatureService) CreateTemperature( temperature models.TemperatureModel ) error {
	query := `INSERT INTO temperature (temperature) VALUES ( ?)`


	_, err := s.DB.Exec(query, temperature.Temperature)

	if err != nil {
		log.Printf("Error al insertar el registro de la temperature: %v", err)
		return err
	}

	fmt.Printf("Registro de temperature insertado correctamente: %v\n", temperature)

	return nil
}