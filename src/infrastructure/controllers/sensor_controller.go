package controllers

import (
	"ApiMulti/src/application"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type SensorController struct {
	service *application.SensorService
}

type SensorRequest struct {
	Estado int `json:"estado"`
}

// New request type for multi-sensor data
type SensorDataRequest struct {
	NumeroSerie        string `json:"numeroSerie"`
	Sensor             string `json:"sensor"`
	FechaActivacion    string `json:"fecha_activacion"`
	FechaDesactivacion string `json:"fecha_desactivacion"`
	Estado             int    `json:"estado"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewSensorController(service *application.SensorService) *SensorController {
	return &SensorController{service: service}
}

// Generate a random serial number
func generateSerialNumber() (string, error) {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Handle registration - new endpoint
func (c *SensorController) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	// Generate a serial number
	serialNumber, err := generateSerialNumber()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error generating serial number"))
		return
	}

	// Return the serial number as plain text
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(serialNumber))
}

// Handle data submission from ESP32 - new endpoint
func (c *SensorController) HandleSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "method not allowed"})
		return
	}

	var req SensorDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid request body"})
		return
	}

	// Process data based on the sensor type
	var err error
	switch req.Sensor {
	case "MQ-2", "MQ2":
		err = c.service.ProcessMQ2Reading(req.Estado)
	case "MQ-135", "MQ135":
		err = c.service.ProcessMQ135Reading(req.Estado)
	case "KY-026", "KY026":
		err = c.service.ProcessKY026Reading(req.Estado)
	case "DHT22":
		err = c.service.ProcessDHT22Reading(req.Estado)
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "unknown sensor type"})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (c *SensorController) HandleKY026(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		sensors, err := c.service.GetAllKY026Readings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sensors)

	case http.MethodPost:
		var req SensorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid request body"})
			return
		}

		if err := c.service.ProcessKY026Reading(req.Estado); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "method not allowed"})
	}
}

func (c *SensorController) HandleMQ2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		sensors, err := c.service.GetAllMQ2Readings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sensors)

	case http.MethodPost:
		var req SensorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid request body"})
			return
		}

		if err := c.service.ProcessMQ2Reading(req.Estado); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "method not allowed"})
	}
}

func (c *SensorController) HandleMQ135(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		sensors, err := c.service.GetAllMQ135Readings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sensors)

	case http.MethodPost:
		var req SensorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid request body"})
			return
		}

		if err := c.service.ProcessMQ135Reading(req.Estado); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "method not allowed"})
	}
}

func (c *SensorController) HandleDHT22(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		sensors, err := c.service.GetAllDHT22Readings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sensors)

	case http.MethodPost:
		var req SensorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid request body"})
			return
		}

		if err := c.service.ProcessDHT22Reading(req.Estado); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "method not allowed"})
	}
}
