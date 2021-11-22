package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func getUrls(w http.ResponseWriter) {
	homeDir := getHomeDir()
	fileName := homeDir + "\\URL_Shortner_db.txt"

	f, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Fprintf(w, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createShortUrls(url string) string {
	// create sha1 hash
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)

	//get first 8 characters only
	shortLinkFull := fmt.Sprintf("%x\n", bs)
	sl := strings.Split(shortLinkFull, "")
	shortLink := "http://localhost:1234/" + strings.Join(sl[:8], "")

	writeToFile(url, shortLink)

	return shortLink
}

func writeToFile(url string, shortLink string) {
	homeDir := getHomeDir()
	fileName := homeDir + "\\URL_Shortner_db.txt"

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writeString := shortLink + "_" + url + "\n"

	_, err2 := f.Write([]byte(writeString))
	if err2 != nil {
		log.Fatal(err2)
	}
}

func searchOriginalUrl(shortLink string) string {
	homeDir := getHomeDir()
	fileName := homeDir + "\\URL_Shortner_db.txt"

	f, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "_")
		if arr[0] == shortLink {
			return arr[1]
		}
	}
	return "Link not found"
}

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}
