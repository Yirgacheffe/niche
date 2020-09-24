package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
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

const tusFolderName = "tus_file_server"

func makeFileDirectory() (string, error) {

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to get current user home directory.", err)
		return "", err
	}

	tusPath := path.Join(userHome, tusFolderName)
	err = os.Mkdir(tusPath, 0744)
	if err != nil {
		log.Println("Unable to create file directory.", err)
	}

	return tusPath, nil

}

// DetailsHandler - /api/files/{id}
func (h *FileHandler) DetailsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)
	fID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, err := h.fileRepo.GetByID(int64(fID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println("Will set Upload-Offset to output.")

	w.Header().Set("Upload-Offset", strconv.Itoa(f.Offset))
	w.WriteHeader(http.StatusOK)

}

// PatchFileHandler - /api/files/{id}
func (h *FileHandler) PatchFileHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, err := h.fileRepo.GetByID(int64(fID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// if file upload complete
	if f.IsComplete == "Y" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Offset validation
	offset, err := strconv.Atoi(r.Header.Get("Upload-Offset"))
	if err != nil {
		log.Println("Wrong upload offset", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Upload offset: %d\n", offset)

	if f.Offset != offset {
		e := fmt.Sprintf("Expect offset: %d, got: %d", f.Offset, offset)
		log.Println(e)

		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(e))
		return
	}

	// Content length validation
	lenInHeader := r.Header.Get("Content-Length")

	cl, err := strconv.Atoi(lenInHeader)
	if err != nil {
		log.Println("Unknown content length in header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	remainLen := f.UploadLength - f.Offset
	if cl != remainLen {
		e := fmt.Sprintf("Content length is not match, expect: %d, got: %d", remainLen, cl)
		log.Println(e)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(e))
		return
	}

	// Write file on the disk, get offset and update the record in database
	// Fake increment the offset length

	dirPath := "/Users/aaron/tus_file_server"

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Receive file partially %s\n", err)
	}

	fp := fmt.Sprintf("%s/%d", dirPath, f.ID)
	savedFile, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Unable to open file %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer savedFile.Close()

	n, err := savedFile.WriteAt(body, int64(offset))
	if err != nil {
		log.Printf("Unable to write file %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Number of bytes written", n)
	f.Offset += n

	uo := strconv.Itoa(f.Offset)
	w.Header().Set("Upload-Offset", uo)

	if f.Offset == f.UploadLength {
		f.IsComplete = "Y"
		log.Println("Uploaded complete successfully!")
	}

	_, err = h.fileRepo.Update(&f)
	if err != nil {
		log.Println("Error happened while updating the file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Going to send succesfully to the response as everything goes fine.")
	w.WriteHeader(http.StatusNoContent)

}

// HealthHandler - Check if it is alive
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	io.WriteString(w, `{"alive": true}`)
}

/*
func responseWithErrorCode(w http.ResponseWriter, code int, message string) {

}
*/
