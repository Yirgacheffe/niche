package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Config struct {
	BackColor   string `json:"BackColor"`
	AppVersion  string `json:"AppVersion"`
	KubePodName string `json:"KubePodName"`
}

type Configs []Config

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {

	rand.Seed(time.Now().UnixNano())
	rdn := rand.Intn(30)

	if rdn <= 20 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {

		var appVersion = os.Getenv("IMAGE_TAG")
		var backColor = "#483D8B"
		var kubePodName = os.Getenv("KUBE_POD_NAME")

		if len(appVersion) == 0 {
			appVersion = "master-testing"
		}

		configs := Config{
			BackColor: backColor, AppVersion: appVersion, KubePodName: kubePodName,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(configs); err != nil {
			panic(err)
		}
	}

}
