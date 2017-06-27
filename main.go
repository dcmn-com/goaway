package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", EchoHandler)
	log.Fatal(http.ListenAndServe(":8765", nil))
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, r.Body)
}
