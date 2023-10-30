package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	DOB       string    `json:"dob" bson:"dob"`
	Gender    string    `json:"gender" bson:"gender"`
	Phone     string    `json:"phone" bson:"phone"`
	NIN       string    `json:"nin" bson:"nin"`
	Pic       string    `json:"pic" bson:"pic"`
	UserType  string    `json:"userType" bson:"userType"`
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

func (*User) FromJSON(user []byte) (User, error) {
	newUser := User{}
	return newUser, json.Unmarshal(user, &newUser)
}
func (*User) ToJSON(user *User) ([]byte, error) {
	return json.Marshal(user)
}
