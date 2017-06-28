package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.HandleFunc("/work", WorkHandler)
	log.Fatal(http.ListenAndServe(":8765", nil))
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, r.Body)
}

func work() {
	time.Sleep(50*time.Millisecond)
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	work()
	EchoHandler(w, r)
}
