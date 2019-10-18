package main

import (
	"encoding/json"
	"fmt"
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

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "RUNNING")
}

func configHandler(w http.ResponseWriter, r *http.Request) {

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

	if backColor == "crimson" {
		r := random(50, 100)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}

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

func testHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/html")
	resp.WriteHeader(http.StatusOK)

	fmt.Fprint(resp, "Running with OK status, Have Fun!")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
