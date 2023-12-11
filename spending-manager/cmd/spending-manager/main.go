package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/spetr/system/test", handleTest)
	http.ListenAndServe(":8080", nil)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  200,
		"message": "Application is working successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
