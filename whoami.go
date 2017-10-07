package main

import (
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	hostname += "\n"

	response := []byte(hostname)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(response)
	})

	http.ListenAndServe(":8080", nil)
}
