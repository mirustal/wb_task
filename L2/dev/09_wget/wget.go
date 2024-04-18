package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pborman/getopt"
)

func getURLandFilename() (string, string) {
	urlPath := getopt.StringLong("url", 'u', "", "URL to download")
	getopt.Parse()
	parsedURL, err := url.ParseRequestURI(*urlPath)
	if err != nil {
		log.Fatalf("Invalid URL: %v", err)
	}

	filename := filepath.Base(parsedURL.Path)
	if filename == "" || strings.HasSuffix(parsedURL.Path, "/") {
		filename = "index.html"
	}
	return *urlPath, filename
}

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Cannot create file: %v", err)
	}
	return file
}

func getData(urlPath string, client *http.Client, file *os.File) int64 {
	resp, err := client.Get(urlPath)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("Error saving data to file: %v", err)
	}
	return size
}

func main() {
	urlPath, filename := getURLandFilename()
	file := createFile(filename)
	defer file.Close()

	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("stopped after 10 redirects")
			}
			return nil
		},
	}

	size := getData(urlPath, client, file)
	fmt.Printf("Downloaded a file %s with size %d bytes\n", filename, size)
}
