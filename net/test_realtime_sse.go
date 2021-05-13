package main

import (
	"encoding/json"
	"fmt"
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

	name := r.FormValue("name")
	message := r.FormValue("message")

	reqMap := map[string]string{"name": name, "message": message}
	json, _ := json.Marshal(reqMap)

	go func() {
		for msgCh := range msgChs {
			msgCh <- []byte(json)
		}
	}()

	w.Write([]byte("ok."))

}

func listenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	msgCh := make(chan []byte)
	msgChs[msgCh] = true

	for {
		select {
		case msg := <-msgCh:
			w.Write(formatSSE("message", string(msg)))
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			delete(msgChs, msgCh)
			fmt.Println("Get Done message from context.")
			return
		}
	}

	// end of this program, nothing to do well this is nouces ......

}

func main() {
	http.HandleFunc("/say", sayHandler)
	http.HandleFunc("/listen", listenHandler)

	log.Println("Running at :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
