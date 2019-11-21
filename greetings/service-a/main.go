package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

type Greeting struct {
	ID          string    `json:"id,omitempty"`
	ServiceName string    `json:"service,omitempty"`
	Message     string    `json:"message,omitempty"`
	CreatedAt   time.Time `json:"created,omitempty"`
}

var greetings []Greeting

func PingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	log.Debug(r)

	greetings = nil

	CallNextServiceWithTrace("http://niche-greetings-b:8080/api/ping", w, r)
	CallNextServiceWithTrace("http://niche-greetings-c:8080/api/ping", w, r)

	tmpGreeting := Greeting{
		ID:          uuid.New().String(),
		ServiceName: "Service-Greetings",
		Message:     "Hello, from Service Greetings!",
		CreatedAt:   time.Now().Local(),
	}

	greetings = append(greetings, tmpGreeting)

	err := json.NewEncoder(w).Encode(greetings)
	if err != nil {
		log.Error(err)
	}

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, err := w.Write([]byte("{\"alive\": true}"))
	if err != nil {
		log.Error(err)
	}

}

func ResponseStatusHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	statusCode, err := strconv.Atoi(params["code"])
	if err != nil {
		log.Error(err)
	}

	w.WriteHeader(statusCode)

}

func CallNextServiceWithTrace(url string, w http.ResponseWriter, r *http.Request) {

	var tmpGreetings []Greeting

	w.Header().Set("Content-Type", "application/json:charset=utf-8")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
	}

	// Http b3 header for Jaeger tracing
	b3Headers := []string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"b3",
	}

	for _, b3Header := range b3Headers {
		if r.Header.Get(b3Header) != "" {
			req.Header.Add(b3Header, r.Header.Get(b3Header))
		}
	}

	log.Info(req)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(body, &tmpGreetings)
	if err != nil {
		log.Error(err)
	}

	for _, r := range tmpGreetings {
		greetings = append(greetings, r)
	}

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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, AllowCredentials: true, AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	})

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/ping", PingHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/health", HealthCheckHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/status/{code}", ResponseStatusHandler).Methods("GET", "OPTIONS")

	api.Handle("/metrics", promhttp.Handler())

	handler := c.Handler(router)
	log.Info(http.ListenAndServe(":8080", handler))

}
