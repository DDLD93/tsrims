package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ddld93/tsrims/auth/controllers"
	"github.com/ddld93/tsrims/auth/models"
	"github.com/ddld93/tsrims/auth/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// response := models.CustomResponse{}
	js, err := utils.GetReqJson(r)
	if err != nil {
		response := models.FormatResponse(false, nil, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	user := models.User{}
	obj, err := user.FromJSON(js)
	if err != nil {
		response := models.FormatResponse(false, nil, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	payload, token, err := controllers.NewCustomerController().Login(obj.Email, obj.Password)
	if err != nil {
		response := models.FormatResponse(false, nil, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	data := map[string]interface{}{
		"user":  payload,
		"token": token,
	}
	response := models.FormatResponse(true, data, "login success")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}
