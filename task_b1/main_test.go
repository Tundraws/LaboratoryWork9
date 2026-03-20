package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected uint64
	}{
		{0, 0}, {1, 1}, {5, 5}, {10, 55},
	}
	for _, tt := range tests {
		if res := calculateFibonacci(tt.input); res != tt.expected {
			t.Errorf("Fibonacci(%d): expected %d, got %d", tt.input, tt.expected, res)
		}
	}
}

func TestHandleCompute(t *testing.T) {
	reqBody, _ := json.Marshal(ComputationRequest{Value: 5})
	req, err := http.NewRequest("POST", "/compute", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCompute)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response ComputationResponse
	json.NewDecoder(rr.Body).Decode(&response)
	if response.Result != 5 {
		t.Errorf("handler returned unexpected result: got %v want %v", response.Result, 5)
	}
}

func TestHandleComputeInvalidMethod(t *testing.T) {
	req, _ := http.NewRequest("GET", "/compute", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCompute)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}