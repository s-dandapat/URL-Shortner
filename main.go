package main

import (
	"fmt"
	"net/http"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	fmt.Println(url)
	id := strings.Replace(url, "/", "", -1)
	shortLink := "http://localhost:1234/" + id
	orginalUrl := searchOriginalUrl(shortLink)
	fmt.Fprintf(w, orginalUrl)
}

func shortUrls(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getUrls(w)

	case "POST":
		url := r.Header.Get("url")
		shortUrl := createShortUrls(url)
		fmt.Fprintf(w, shortUrl)

	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported !")
	}
}

func main() {
	fmt.Println("Listening to Short URL...")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/url", shortUrls)

	http.ListenAndServe(":1234", nil)
}
