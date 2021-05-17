package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	rawHtml := string(body)

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	rawHtml = re.ReplaceAllStringFunc(rawHtml, strings.ToLower)

	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	rawHtml = re.ReplaceAllString(rawHtml, "")

	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	rawHtml = re.ReplaceAllString(rawHtml, "")

	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	rawHtml = re.ReplaceAllString(rawHtml, "\n")

	re, _ = regexp.Compile("\\s{2,}")
	rawHtml = re.ReplaceAllString(rawHtml, "\n")

	fmt.Println(strings.TrimSpace(rawHtml))

}
