package handlers

import (
	"net/http"

	"github.com/farbautie/ygg/internal/services"
)

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(service *services.FileService) *FileHandler {
	return &FileHandler{
		service: service,
	}
}

func (fh *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := fh.service.Save(handler.Filename, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true, "message": "File uploaded successfully"}`))
}
