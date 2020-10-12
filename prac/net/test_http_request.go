package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	reqURL, _ := url.Parse("http://dummy.restapiexample.com/api/v1/create")

	reqBody := ioutil.NopCloser(strings.NewReader(`
		{
			"name": "Yuan",
			"age": 45,
			"salary": 5000
		}
	`))

	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{"Content-Type": {"application-type; charset=UTF-8"}},
		Body:   reqBody,
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println("Response Status:", res.StatusCode)
	fmt.Println(data)

}
