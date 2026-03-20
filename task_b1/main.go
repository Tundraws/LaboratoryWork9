package main

import (
	"encoding/json"
	"net/http"
)

const BindAddress = ":8080"

type ComputationRequest struct {
	Value int `json:"value"`
}

type ComputationResponse struct {
	Result uint64 `json:"result"`
	Status string `json:"status"`
}

func calculateFibonacci(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	return calculateFibonacci(n-1) + calculateFibonacci(n-2)
}

func handleCompute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req ComputationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := ComputationResponse{
		Result: calculateFibonacci(req.Value),
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/compute", handleCompute)
	http.ListenAndServe(BindAddress, nil)
}