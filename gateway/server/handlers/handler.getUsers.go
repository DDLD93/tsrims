package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ddld93/tsrims/auth/controllers"
	"github.com/ddld93/tsrims/auth/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	payload, err := controllers.NewCustomerController().GetAllUsers()
	if err != nil {
		response := models.FormatResponse(false, nil, err.Error());
		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(response);
		return
	}
	response := models.FormatResponse(true, payload, "");

	w.WriteHeader(http.StatusOK);
	json.NewEncoder(w).Encode(response);

}
