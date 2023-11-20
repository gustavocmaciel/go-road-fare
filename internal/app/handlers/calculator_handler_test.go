package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetVehicleTax(t *testing.T) {
	// Create a new request with a sample URL
	req, err := http.NewRequest("GET", "/api/calculator/car/2023-05-05T12:00:00,2023-05-05T18:15:00", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new instance of the router
	r := mux.NewRouter()
	r.HandleFunc("/api/calculator/{vehicle}/{time}", GetVehicleTax)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, rr.Code)
	}

	// Parse the response body
	var tax int
	if err := json.Unmarshal(rr.Body.Bytes(), &tax); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Perform assertions on the result
	expectedTax := 8
	if tax != expectedTax {
		t.Errorf("Expected tax %v, but got %v", expectedTax, tax)
	}
}
