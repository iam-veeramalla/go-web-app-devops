// Test the main function

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Just verify the code not html content
	expected := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expected)
	}
}
