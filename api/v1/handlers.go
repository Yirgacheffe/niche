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
	Key          string `json:"Key"`
	BackColor    string `json:"BackColor"`
	AppVersion   string `json:"AppVersion"`
	BuildDate    string `json:"BuildDate"`
	KubeNodeName string `json:"KubeNodeName"`
	KubePodName  string `json:"KubePodName"`
	KubePodIP    string `json:"KubePodIP"`
}

type Configs []Config

func ConfigHandler(w http.ResponseWriter, r *http.Request) {

	var appVersion = os.Getenv("IMAGE_TAG")
	var backColor = "asparagus"
	var imageBuildeDate = os.Getenv("IMAGE_BUILD_DATE")

	var kubeNodeName = os.Getenv("KUBE_NODE_NAME")
	var kubePodName = os.Getenv("KUBE_POD_NAME")
	var kubePodIP = os.Getenv("KUBE_POD_IP")

	if len(appVersion) == 0 {
		appVersion = "master-testing"
	}

	configs := Config{Key: "10", BackColor: backColor, AppVersion: appVersion, BuildDate: imageBuildeDate, KubeNodeName: kubeNodeName, KubePodName: kubePodName, KubePodIP: kubePodIP}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(configs); err != nil {
		panic(err)
	}

}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
