package middlewares

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddld93/incident/models"
	"github.com/google/uuid"
)

// FileUploadMiddleware handles file uploads and chains with other handlers.
func FileUploadMiddleware(fieldName string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost && strings.Split(r.Header.Get("Content-Type"), ";")[0] == "multipart/form-data" {
            file, header, err := r.FormFile(fieldName)
            if err != nil {
                response := models.FormatResponse(false, nil, "File upload is required")
                jsonResponse, _ := json.Marshal(response)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusBadRequest)
                _, _ = w.Write(jsonResponse)
                return
            }
            defer file.Close()

            extension := filepath.Ext(header.Filename)
            uuidString := strings.ReplaceAll(uuid.New().String(), "-", "")
            uuidFilename := uuidString + extension
            filename := filepath.Join("./uploads", uuidFilename)

            outFile, err := os.Create(filename)
            if err != nil {
                response := models.FormatResponse(false, nil, "Failed to create file on server")
                jsonResponse, _ := json.Marshal(response)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                _, _ = w.Write(jsonResponse)
                return
            }
            defer outFile.Close()

            _, err = io.Copy(outFile, file)
            if err != nil {
                response := models.FormatResponse(false, nil, "Failed to save file on server")
                jsonResponse, _ := json.Marshal(response)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                _, _ = w.Write(jsonResponse)
                return
            }

            r = r.WithContext(context.WithValue(r.Context(), "filepath", filename))
        }else{
            response := models.FormatResponse(false, nil, "inavalid content-type")
            jsonResponse, _ := json.Marshal(response)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            _, _ = w.Write(jsonResponse)
            return
        }

        // Call the next handler in the chain.
        next(w, r)
    }
}
