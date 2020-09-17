package main

import (
	"fmt"
	"net/http"
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

// GetByID ... handle details operation
func (h FileHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WIP")
}
