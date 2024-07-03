package models

type Response struct {
	Status  int64       `json:"Status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
