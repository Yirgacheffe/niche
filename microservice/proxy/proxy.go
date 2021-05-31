package main

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

type Proxy struct {
	Client  *http.Client
	BaseURL string
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		log.Printf("error occurred during process request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := p.Client.Do(r)
	if err != nil {
		log.Printf("error occurred during client operation: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	CopyResponse(w, resp) // Copy target 'resp' to current writer
}

func (p *Proxy) ProcessRequest(r *http.Request) error {
	rawURL := p.BaseURL + r.URL.String()

	proxyURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	r.Host = proxyURL.Host
	r.URL = proxyURL
	r.RequestURI = ""

	return nil
}

func CopyResponse(w http.ResponseWriter, resp *http.Response) {
	var out bytes.Buffer
	out.ReadFrom(resp.Body)

	// Write headers
	for k, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(out.Bytes()) // Writes everything to ResponseWriter
}
