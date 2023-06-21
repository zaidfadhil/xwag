package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("unexpected Content-Type header: got %s, want text/html; charset=utf-8", contentType)
	}

	expectedHTML := fmt.Sprintf(htmlindex, "swagger.yaml")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body[:len(expectedHTML)]) != expectedHTML {
		t.Errorf("unexpected response body: got %s, want %s", body, expectedHTML)
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
	if contentType != "application/yaml" {
		t.Errorf("unexpected Content-Type header: got %s, want application/yaml", contentType)
	}

	swaggerFile, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != string(swaggerFile) {
		t.Errorf("unexpected response body: got %s, want %s", body, swaggerFile)
	}
}
