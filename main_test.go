package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSwaggerHandler(t *testing.T) {
	file = "testdata/test.yaml"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	swaggerHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	contentType := res.Header().Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if contentType != expectedContentType {
		t.Errorf("unexpected Content-Type header: got %s, want %s", contentType, expectedContentType)
	}

	expectedHTML := fmt.Sprintf(htmlTemplate, "swagger.yaml")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expectedHTML {
		t.Errorf("unexpected response body:\nGot:\n%s\n\nWant:\n%s", string(body), expectedHTML)
	}
}

func TestSwaggerYAMLHandler(t *testing.T) {
	file = "testdata/test.yaml"

	req, err := http.NewRequest("GET", "/swagger.yaml", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	swaggerYAMLHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	contentType := res.Header().Get("Content-Type")
	expectedContentType := "application/yaml"
	if contentType != expectedContentType {
		t.Errorf("unexpected Content-Type header: got %s, want %s", contentType, expectedContentType)
	}

	swaggerFile, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != string(swaggerFile) {
		t.Errorf("unexpected response body:\nGot:\n%s\n\nWant:\n%s", string(body), string(swaggerFile))
	}
}
