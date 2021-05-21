package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "find_links2: %v", err)
			continue
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}

}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("find link: %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links // visit & appends to links ......
}

func findLinksWithLog(url string) ([]string, error) {
	log.Printf(
		"findLinks %s", url,
	)
	return findLinks(url)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return // 'err' returned by default
	}

	defer resp.Body.Close()

	doc, err = html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML:%s", err)
		return
	}

	// How to count the words and images ?
	words, images := countWordsAndImages(doc) // TODO: Not finished this func
	return
}
