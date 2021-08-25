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
	htmlHeader    = "<!DOCTYPE html><html><head><style>table, th, td {border: 1px solid black;font-family: 'Courier New';font-size: 20px;color: white}th, td {padding: 10px;}</style></head><font color=black><h1>Istio Canary Demo Homepage - 2021</h1><body style=background-color:white>"
	htmlTitle     = "<p>Repo Git: %s <br>Web image build date: %s <br>Running on: (%s / %s)</p><br><table>"
	htmlTableCell = "<td bgcolor=%s align=center>%s</td>"
)

var (
	respUnavailable = fmt.Sprintf(htmlTableCell, "#A9A9A9", "Service Unavailable!")
	respIncorrect   = fmt.Sprintf(htmlTableCell, "#A9A9A9", "Incorrect Response!")
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

	makeTitleHtml := func() string {
		var (
			gitSHA         = os.Getenv("GIT_SHA")
			imageBuildDate = os.Getenv("IMAGE_BUILD_DATE")
			kubePodName    = os.Getenv("KUBE_POD_NAME")
			kubePodIP      = os.Getenv("KUBE_POD_IP")
		)

		if len(gitSHA) == 0 {
			gitSHA = "Not set"
		}
		if len(imageBuildDate) == 0 {
			imageBuildDate = "8/25/2021 22:41:31"
		}
		if len(kubePodName) == 0 {
			kubePodName = "niche-web-1659604661-zh6rp"
		}
		if len(kubePodIP) == 0 {
			kubePodIP = "192.168.1.100"
		}

		return fmt.Sprintf(htmlTitle, gitSHA, imageBuildDate, kubePodName, kubePodIP)
	}

	configAPIUrl := func() string {
		var (
			service = os.Getenv("API_SERVICE")
			port    = os.Getenv("API_PORT")
		)

		if len(service) == 0 {
			service = "localhost"
		}
		if len(port) == 0 {
			port = "8081"
		}

		return fmt.Sprintf("http://%s:%s/api/configs", service, port)
	}

	fmt.Fprintf(w, htmlHeader)
	fmt.Fprintf(w, makeTitleHtml())

	// Get Url param 'City', set to Header, ignore others
	apiUrl := configAPIUrl()
	paramCity := r.URL.Query().Get("city")

	// Make channel to call api in go routine
	done := make(chan bool)
	defer close(done)

	maxRoutines := 25
	resultGen := func(done <-chan bool) <-chan string {
		result := make(chan string)

		go func() {
			defer close(result)
			for i := 0; i < maxRoutines; i++ {
				select {
				case <-done:
					return
				case result <- createTableCell(paramCity, apiUrl):
				}
			}
		}()

		return result
	}

	// loop throught the api to build table
	for row := 0; row < 5; row++ {
		fmt.Fprintf(w, "<tr>")
		for col := 0; col < 5; col++ {
			fmt.Fprintf(w, <-resultGen(done))
		}
		fmt.Fprintf(w, "</tr>")
	}

	fmt.Fprintf(w, "</table></body></html>")
}

func createTableCell(paramCity, apiUrl string) string {

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(paramCity) != 0 {
		req.Header.Add("city", paramCity)
	}

	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		return respUnavailable
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respUnavailable
	}

	log.Printf(string(respData))

	var conf Config
	err = json.Unmarshal(respData, &conf)
	if err != nil {
		return respIncorrect
	}

	respCorrect := fmt.Sprintf(htmlTableCell, conf.BgColor, conf.Version+":"+conf.PodName)
	return respCorrect
}
