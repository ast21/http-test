package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := make(map[string]string)
	data["hostname"], err = os.Hostname()
	if err != nil {
		data["hostname"] = "unknown"
	}

	jsonData, _ := json.Marshal(data)

	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
