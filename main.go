package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	workDur    = 50 * time.Millisecond
	numWorkers = 4
)

func main() {
	HandleFunc("/echo", EchoHandler)
	HandleFunc("/work", WorkHandler)
	HandleFunc("/worklimit", WorkLimitHandler)
	HandleFunc("/worklimitshed", ShedLimit(WorkLimitHandler, numWorkers))
	HandleFunc("/worklimitshedb", ShedLimitTimeout(WorkLimitHandler, numWorkers, 5*time.Millisecond))
	log.Fatal(http.ListenAndServe(":8765", nil))
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, r.Body)
}

func work() {
	time.Sleep(workDur)
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	work()
	EchoHandler(w, r)
}

var workers = make(chan struct{}, numWorkers)

func WorkLimitHandler(w http.ResponseWriter, r *http.Request) {
	workers <- struct{}{}
	WorkHandler(w, r)
	<-workers
}
