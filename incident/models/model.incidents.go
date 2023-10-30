package models

import (
	"time"
)

type Incident struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    string    `json:"userId" bson:"userId"`
	Channel   string    `json:"channel" bson:"channel"`
	Phone     string    `json:"phone" bson:"phone"`
	Latitude  float32   `json:"latitude" bson:"latitude"`
	Logitude  float32   `json:"longitude" bson:"longitude"`
	Images    []string  `json:"images" bson:"images"`
	Audio     string    `json:"audio" bson:"audio"`
	Video     string    `json:"video" bson:"video"`
	Type      string    `json:"type" bson:"type"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type customResponse struct {
	Ok      bool        `json:"ok"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func FormatResponse(ok bool, data interface{}, message string) *customResponse {
	response := customResponse{
		Ok:      ok,
		Data:    data,
		Message: message,
	}
	// jsonResponse, _ := json.Marshal(response)
	return &response
}
