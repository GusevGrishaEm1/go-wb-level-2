package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadPage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := "index.html"
	if strings.HasSuffix(url, "/") {
		filename += "index.html"
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	url := "https://www.example.com"
	err := downloadPage(url)
	if err != nil {
		fmt.Println("Error downloading page:", err)
		return
	}
	fmt.Println("Page downloaded successfully!")
}
