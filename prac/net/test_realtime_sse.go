package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func formatSSE(event string, data string) []byte {

	payload := "event: " + event + "\n"

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		payload = payload + "data:" + line + "\n"
	}

	return []byte(payload + "\n")

}

var msgChs = make(map[chan []byte]bool)

func sayHandler(w http.ResponseWriter, r *http.Request) {

	name    := r.FormValue("name")
	message := r.FormValue("message")
	
	reqMap := map[string]string{"name": name, "message": message}
	json, _ := json.Marshal(reqMap)

	go func() {
		for msgCh := range msgChs {
			msgCh <- []byte(json)
		}
	}

	w.Write([]byte("ok."))	// write ok to response
	
}

func listenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func main() {
	http.HandleFunc("/say", sayHandler)
	http.HandleFunc("listen", listenHandler)

	log.Println("Running at :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
