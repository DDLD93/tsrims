package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ddld93/tsrims/auth/controllers"
	"github.com/ddld93/tsrims/auth/models"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	id := params["id"]
	payload, err := controllers.NewCustomerController().GetUserByID(id)
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
