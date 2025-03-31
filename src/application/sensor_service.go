package application

import (
	"ApiMulti/src/core"
	"ApiMulti/src/domain/entities"
	"ApiMulti/src/domain/repositories"
	"encoding/json"
	"log"
	"os"
	"time"
)

type SensorService struct {
	repo repositories.SensorRepository
}

func NewSensorService(repo repositories.SensorRepository) *SensorService {
	return &SensorService{repo: repo}
}

func (s *SensorService) ProcessKY026Reading(estado int) error {
	sensor := &entities.KY026{
		FechaActivacion: time.Now().Format("2006-01-02 15:04:05"),
		Estado:          estado,
	}

	if err := s.repo.SaveKY026(sensor); err != nil {
		log.Printf("Error saving KY026 to database: %v", err)
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		log.Printf("Error marshaling KY026 message: %v", err)
		return err
	}

	queueName := os.Getenv("RABBITMQ_QUEUE_KY026")
	if err := core.PublishMessage(queueName, message); err != nil {
		log.Printf("Error publishing KY026 message: %v", err)
		return err
	}

	log.Printf("Successfully processed KY026 reading: %+v", sensor)
	return nil
}

func (s *SensorService) ProcessMQ2Reading(estado int) error {
	sensor := &entities.MQ2{
		FechaActivacion: time.Now().Format("2006-01-02 15:04:05"),
		Estado:          estado,
	}

	if err := s.repo.SaveMQ2(sensor); err != nil {
		log.Printf("Error saving MQ2 to database: %v", err)
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		log.Printf("Error marshaling MQ2 message: %v", err)
		return err
	}

	queueName := os.Getenv("RABBITMQ_QUEUE_MQ2")
	if err := core.PublishMessage(queueName, message); err != nil {
		log.Printf("Error publishing MQ2 message: %v", err)
		return err
	}

	log.Printf("Successfully processed MQ2 reading: %+v", sensor)
	return nil
}

func (s *SensorService) ProcessMQ135Reading(estado int) error {
	sensor := &entities.MQ135{
		FechaActivacion: time.Now().Format("2006-01-02 15:04:05"),
		Estado:          estado,
	}

	if err := s.repo.SaveMQ135(sensor); err != nil {
		log.Printf("Error saving MQ135 to database: %v", err)
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		log.Printf("Error marshaling MQ135 message: %v", err)
		return err
	}

	queueName := os.Getenv("RABBITMQ_QUEUE_MQ135")
	if err := core.PublishMessage(queueName, message); err != nil {
		log.Printf("Error publishing MQ135 message: %v", err)
		return err
	}

	log.Printf("Successfully processed MQ135 reading: %+v", sensor)
	return nil
}

func (s *SensorService) ProcessDHT22Reading(estado int) error {
	sensor := &entities.DHT22{
		FechaActivacion: time.Now().Format("2006-01-02 15:04:05"),
		Estado:          estado,
	}

	if err := s.repo.SaveDHT22(sensor); err != nil {
		log.Printf("Error saving DHT22 to database: %v", err)
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		log.Printf("Error marshaling DHT22 message: %v", err)
		return err
	}

	queueName := os.Getenv("RABBITMQ_QUEUE_DHT22")
	if err := core.PublishMessage(queueName, message); err != nil {
		log.Printf("Error publishing DHT22 message: %v", err)
		return err
	}

	log.Printf("Successfully processed DHT22 reading: %+v", sensor)
	return nil
}

func (s *SensorService) GetAllKY026Readings() ([]*entities.KY026, error) {
	sensors, err := s.repo.GetAllKY026()
	if err != nil {
		log.Printf("Error getting all KY026 readings: %v", err)
		return nil, err
	}
	return sensors, nil
}

func (s *SensorService) GetAllMQ2Readings() ([]*entities.MQ2, error) {
	sensors, err := s.repo.GetAllMQ2()
	if err != nil {
		log.Printf("Error getting all MQ2 readings: %v", err)
		return nil, err
	}
	return sensors, nil
}

func (s *SensorService) GetAllMQ135Readings() ([]*entities.MQ135, error) {
	sensors, err := s.repo.GetAllMQ135()
	if err != nil {
		log.Printf("Error getting all MQ135 readings: %v", err)
		return nil, err
	}
	return sensors, nil
}

func (s *SensorService) GetAllDHT22Readings() ([]*entities.DHT22, error) {
	sensors, err := s.repo.GetAllDHT22()
	if err != nil {
		log.Printf("Error getting all DHT22 readings: %v", err)
		return nil, err
	}
	return sensors, nil
}
