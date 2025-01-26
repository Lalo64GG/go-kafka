package models

type TemperatureModel struct {
	ID          int    `json:"id"`
	Temperature float64 `json:"temperature"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}