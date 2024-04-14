package main

import (
	"bytes"
	"os"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	url := "https://www.example.com"
	err := downloadPage(url)
	if err != nil {
		t.Errorf("Error downloading page: %v", err)
	}

	// Проверяем, что файл был создан
	if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		t.Errorf("Expected file index.html to be created, but it wasn't")
	}

	// Проверяем содержимое файла
	content, err := os.ReadFile("index.html")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	expectedContent := []byte("Example Domain")
	if !bytes.Contains(content, expectedContent) {
		t.Errorf("Expected content not found in the downloaded page")
	}

	// Чистим за собой после теста
	os.Remove("index.html")
}
