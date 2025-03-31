package repositories

import "ApiMulti/src/domain/entities"

type SensorRepository interface {
	SaveKY026(sensor *entities.KY026) error
	SaveMQ2(sensor *entities.MQ2) error
	SaveMQ135(sensor *entities.MQ135) error
	SaveDHT22(sensor *entities.DHT22) error
	GetKY026ByID(id int) (*entities.KY026, error)
	GetMQ2ByID(id int) (*entities.MQ2, error)
	GetMQ135ByID(id int) (*entities.MQ135, error)
	GetDHT22ByID(id int) (*entities.DHT22, error)
	SaveESP32(esp32 *entities.ESP32) error
	GetAllKY026() ([]*entities.KY026, error)
	GetAllMQ2() ([]*entities.MQ2, error)
	GetAllMQ135() ([]*entities.MQ135, error)
	GetAllDHT22() ([]*entities.DHT22, error)
}
