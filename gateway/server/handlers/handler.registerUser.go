package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ddld93/tsrims/auth/controllers"
	"github.com/ddld93/tsrims/auth/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := models.FormatResponse(false, nil, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	user.CreatedAt = time.Now()
	payload, err := controllers.NewCustomerController().Register(&user)
	if err != nil {
		response := models.FormatResponse(false, nil, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	payload.Password = "******"
	response := models.FormatResponse(true, payload, "")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
