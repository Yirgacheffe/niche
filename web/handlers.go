package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	requestID            = "x-request-id"
	zipkinB3TraceID      = "x-b3-traceid"
	zipkinB3SpanID       = "x-b3-spanid"
	zipkinB3ParentSpanID = "x-b3-parentspanid"
	zipkinB3Sampled      = "x-b3-sampled"
	zipkinB3Flags        = "x-b3-flags"
	zipkinB3             = "b3"
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

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var gitSHA = os.Getenv("GIT_SHA")
	if len(gitSHA) == 0 {
		gitSHA = "Not set"
	}

	var imageBuildDate = os.Getenv("IMAGE_BUILD_DATE")
	if len(imageBuildDate) == 0 {
		imageBuildDate = "9/23/2019 22:41:31"
	}

	var kubePodName = os.Getenv("KUBE_POD_NAME")
	if len(kubePodName) == 0 {
		kubePodName = "niche-web-1659604661-zh6rp"
	}

	var kubePodIP = os.Getenv("KUBE_POD_IP")
	if len(kubePodIP) == 0 {
		kubePodIP = "192.168.1.100"
	}

	var htmlHeader = "<!DOCTYPE html><html><head><style>table, th, td {border: 1px solid black;font-family: 'Courier New';font-size: 28px;color: white}th, td {padding: 20px;}</style></head><font color=black><h1>Istio Canary Demo Homepage - 2019</h1><body style=background-color:white>"

	fmt.Fprintf(w, htmlHeader)
	fmt.Fprintf(w, "<p>Repo Git: %s <br>Web image build date: %s <br>Running on: (%s / %s)</p><br><table>", gitSHA, imageBuildDate, kubePodName, kubePodIP)

	// loop throught the api to build table
	i := 1

	for i <= 5 {
		fmt.Fprintf(w, "<tr>")

		j := 1
		for j <= 5 {
			fmt.Fprintf(w, createTableCell())
			j = j + 1
		}

		fmt.Fprintf(w, "</tr>")
		i = i + 1
	}

	// render html footer
	fmt.Fprintf(w, "</table></body></html>")

}

func testHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/html")
	resp.WriteHeader(http.StatusOK)

	fmt.Fprint(resp, "Running with OK status, Have Fun!")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{ "alive": true }`)
}

func createTableCell() string {

	var apiService = os.Getenv("API_SERVICE")
	if len(apiService) == 0 {
		apiService = "localhost"
	}

	var apiPort = os.Getenv("API_PORT")
	if len(apiPort) == 0 {
		apiPort = "80"
	}

	url := "http://" + apiService + ":" + apiPort + "/configs"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(string(responseData))

	var configObj Config
	json.Unmarshal(responseData, &configObj)
	backColor := configObj.BackColor
	apiVersion := configObj.AppVersion
	podName := configObj.KubePodName

	return "<td bgcolor=" + backColor + " align=center>" + apiVersion + ":" + podName + "</td>"

}

func searchHandler(w http.ResponseWriter, r *http.Request) {

	var gitSHA = os.Getenv("GIT_SHA")
	if len(gitSHA) == 0 {
		gitSHA = "Not set"
	}

	var imageBuildDate = os.Getenv("IMAGE_BUILD_DATE")
	if len(imageBuildDate) == 0 {
		imageBuildDate = "9/23/2019 22:41:31"
	}

	var kubePodName = os.Getenv("KUBE_POD_NAME")
	if len(kubePodName) == 0 {
		kubePodName = "niche-web-1659604661-zh6rp"
	}

	var kubePodIP = os.Getenv("KUBE_POD_IP")
	if len(kubePodIP) == 0 {
		kubePodIP = "192.168.1.100"
	}

	var htmlHeader = "<!DOCTYPE html><html><head><style>table, th, td {border: 1px solid black;font-family: 'Courier New';font-size: 28px;color: white}th, td {padding: 20px;}</style></head><font color=black><h1>Istio Canary Demo Homepage - 2019</h1><body style=background-color:white>"

	fmt.Fprintf(w, htmlHeader)
	fmt.Fprintf(w, "<p>Repo Git: %s <br>Web image build date: %s <br>Running on: (%s / %s)</p><br><table>", gitSHA, imageBuildDate, kubePodName, kubePodIP)

	// loop throught the api to build table
	i := 1

	for i <= 5 {
		fmt.Fprintf(w, "<tr>")

		j := 1
		for j <= 5 {
			fmt.Fprintf(w, createTableCellWithCityAndTracing(r))
			j = j + 1
		}

		fmt.Fprintf(w, "</tr>")
		i = i + 1
	}

	// render html footer
	fmt.Fprintf(w, "</table></body></html>")

}

func createTableCellWithCityAndTracing(r *http.Request) string {

	var apiService = os.Getenv("API_SERVICE")
	if len(apiService) == 0 {
		apiService = "localhost"
	}

	var apiPort = os.Getenv("API_PORT")
	if len(apiPort) == 0 {
		apiPort = "80"
	}

	url := "http://" + apiService + ":" + apiPort + "/configs"

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get Url param 'City', set to Header
	params := r.URL.Query()
	city := params.Get("city")

	if len(city) != 0 {
		request.Header.Add("city", city)
	}

	// Headers must be passed for Jaeger Distributed Tracing
	b3Headers := []string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanId",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"b3",
	}

	for _, b3Header := range b3Headers {
		if r.Header.Get(b3Header) != "" {
			request.Header.Add(b3Header, r.Header.Get(b3Header))
		}
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(string(responseData))

	var configObj Config
	json.Unmarshal(responseData, &configObj)
	backColor := configObj.BackColor
	apiVersion := configObj.AppVersion
	podName := configObj.KubePodName

	return "<td bgcolor=" + backColor + "align=center>" + apiVersion + ":" + podName + "</td>"

}
