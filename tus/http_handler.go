package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NewFileHandler ... create http handler for file op
func NewFileHandler(db *DB) *FileHandler {

	return &FileHandler{
		fileRepo: NewMySQLFileRepo(db.SQL),
	}

}

// FileHandler ... Http handler
type FileHandler struct {
	fileRepo FileRepo
}

// DetailsHandler - /api/files/{id}
func (h *FileHandler) DetailsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)
	fID, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, err := h.fileRepo.GetByID(int64(fID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	log.Println("Will set Upload-Offset to output.")

	w.Header().Set("Upload-Offset", strconv.Itoa(f.Offset))
	w.WriteHeader(http.StatusOK)

}

// HealthHandler - Check if it is alive
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	io.WriteString(w, `{"alive": true}`)
}
