package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Set up a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function
	handler := http.HandlerFunc(helloHandler)

	// Set the environment variable
	os.Setenv("NAME", "Test")

	// Call the handler function
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Hello World from Test!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
