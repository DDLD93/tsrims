package handlers

// import (
// 	"net/http"

// 	"github.com/ddld93/incident/models"
// )
	

// func AddIncidentVideo(w http.ResponseWriter, r *http.Request) {
//     var incident models.Incident
//     err := json.NewDecoder(r.Body).Decode(&incident)
//     if err != nil {
//         response := models.FormatResponse(false, nil, err.Error())
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(response)
//         return
//     }

//     videoURL := "" 
//     incident.Video = videoURL

//     controller := controllers.NewCustomerController()
//     addedIncident, err := controller.AddIncident(&incident)
//     if err != nil {
//         response := models.FormatResponse(false, nil, err.Error())
//         w.WriteHeader(http.StatusInternalServerError)
//         json.NewEncoder(w).Encode(response)
//         return
//     }

//     response := models.FormatResponse(true, addedIncident, "Incident with video added successfully")
//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(response)
// }

