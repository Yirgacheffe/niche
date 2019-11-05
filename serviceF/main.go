package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

// Greeting message transparent between service in json format
type Greeting struct {
	ID          string    `json:"id,omitempty"`
	ServiceName string    `json:"service,omitempty"`
	Message     string    `json:"message,omitempty"`
	CreatedAt   time.Time `json:"created,omitempty"`
}

var greetings []Greeting

func PingHandler() {

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json; charset=utf-8")

	_, err := w.Write([]byte("{\"alive\": true}"))
	if err != nil {
		log.Error(err)
	}

}

func deserialization(b []byte) Greeting {

	log.Debug(b)

	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)

	var tmpGreeting Greeting
	err := decoder.Decode(&tmpGreeting)
	if err != nil {
		log.Error(err)
	}

	return tmpGreeting

}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback

}

func init() {

	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true

	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(getEnv("LOG_LEVEL", "info"))
	if err != nil {
		log.Error(err)
	}

	log.SetLevel(level)

}

func main() {

	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true

	log.SetFormatter(&formatter)

}
