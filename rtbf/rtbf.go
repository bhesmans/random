package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fetch(os.Args[1])
}

// Code sample from goquery example.
func findInUrl(url, filter, attr string) string {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	val, ok := doc.Find(filter).Attr(attr)

	if !ok {
		log.Fatalf("Can't find %s in %s", filter, url)
	}
	return val
}

func fetch(url string) {

	frame := findInUrl(url, "iframe", "src")

	dataMedia := findInUrl(frame, "div#js-embed-player", "data-media")

	fmt.Println(dataMedia)
}
