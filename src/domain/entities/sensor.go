package entities

type KY026 struct {
	ID                 int    `json:"id"`
	FechaActivacion    string `json:"fecha_activacion"`
	FechaDesactivacion string `json:"fecha_desactivacion"`
	Estado             int    `json:"estado"`
}

type MQ2 struct {
	ID                 int    `json:"id"`
	FechaActivacion    string `json:"fecha_activacion"`
	FechaDesactivacion string `json:"fecha_desactivacion"`
	Estado             int    `json:"estado"`
}

type MQ135 struct {
	ID                 int    `json:"id"`
	FechaActivacion    string `json:"fecha_activacion"`
	FechaDesactivacion string `json:"fecha_desactivacion"`
	Estado             int    `json:"estado"`
}

type DHT22 struct {
	ID                 int    `json:"id"`
	FechaActivacion    string `json:"fecha_activacion"`
	FechaDesactivacion string `json:"fecha_desactivacion"`
	Estado             int    `json:"estado"`
}

type ESP32 struct {
	ID      int   `json:"id"`
	KY026ID int   `json:"ky026_id"`
	MQ2ID   int   `json:"mq2_id"`
	MQ135ID int   `json:"mq135_id"`
	DHT22ID int   `json:"dht22_id"`
	KY026   KY026 `json:"ky026"`
	MQ2     MQ2   `json:"mq2"`
	MQ135   MQ135 `json:"mq135"`
	DHT22   DHT22 `json:"dht22"`
}
