package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	htmlHeader = "<!DOCTYPE html><html><head><style>table, th, td {border: 1px solid black;font-family: 'Courier New';font-size: 20px;color: white}th, td {padding: 10px;}</style></head><font color=black><h1>Istio Canary Demo Homepage - 2019</h1><body style=background-color:white>"
	htmlTitle  = "<p>Repo Git: %s <br>Web image build date: %s <br>Running on: (%s / %s)</p><br><table>"
)

type Config struct {
	BgColor string `json:"bg_color,omitempty"`
	Version string `json:"version,omitempty"`
	PodName string `json:"pod_name,omitempty"`
}

// HomeHandler handle Index page, just say hello
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./tmpls/index.html"))

	// Just practice the template
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Hi, Seems an error happened to render Index Page.")
	}

}

// HealthCheckHandler is necessary for container to check liveness
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{ "alive": true }`)
}

// SearchHandler handle search endpoint from the client
func SearchHandler(w http.ResponseWriter, r *http.Request) {

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

	fmt.Fprintf(w, htmlHeader)
	fmt.Fprintf(w, htmlTitle, gitSHA, imageBuildDate, kubePodName, kubePodIP)

	// Make channel to call api in go routine
	ch := make(chan string)
	i := 1
	for i <= 25 {
		go createTableCell(r, ch)
		i = i + 1
	}

	// loop throught the api to build table
	i = 1

	for i <= 5 {
		fmt.Fprintf(w, "<tr>")

		j := 1
		for j <= 5 {
			fmt.Fprintf(w, <-ch)
			j = j + 1
		}

		fmt.Fprintf(w, "</tr>")
		i = i + 1
	}

	fmt.Fprintf(w, "</table></body></html>")

}

func createTableCell(r *http.Request, ch chan<- string) {

	var apiService = os.Getenv("API_SERVICE")
	if len(apiService) == 0 {
		apiService = "localhost"
	}

	var apiPort = os.Getenv("API_PORT")
	if len(apiPort) == 0 {
		apiPort = "8081"
	}

	url := "http://" + apiService + ":" + apiPort + "/api/configs"

	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get Url param 'City', set to Header, ignore others
	params := r.URL.Query()
	city := params.Get("city")

	if len(city) != 0 {
		request.Header.Add("city", city)
	}
	// end of the headers

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		ch <- "<td bgcolor=#A9A9A9 align=center>Service Unavailable!</td>"
		return
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		ch <- "<td bgcolor=#A9A9A9 align=center>Service Unavailable!</td>"
		return
	}

	log.Printf(string(responseData))

	var configObj Config
	err = json.Unmarshal(responseData, &configObj)
	if err != nil {
		log.Println(err)
		ch <- "<td bgcolor=#A9A9A9 align=center>Incorrect Response!</td>"
		return
	}

	backColor := configObj.BgColor
	apiVersion := configObj.Version
	podName := configObj.PodName

	ch <- "<td bgcolor=" + backColor + " align=center>" + apiVersion + ":" + podName + "</td>"
	return

}
