package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"golang.org/x/crypto/bcrypt"
)

func GuessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	msg := r.FormValue("message")
	real := []byte("$2a$10$2ovnPWuIjMx2S0HvCxP/mutzdsGhyt8rq/JqnJg/6OyC3B0APMGlK")

	if err := bcrypt.CompareHashAndPassword(real, []byte(msg)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Nope!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Right!"))
	return
}

func main() {
	http.HandleFunc("/guess", GuessHandler)

	fmt.Println("server started at localhost:8080")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}
