package main

import (
	"net/http"
	"time"
)

func ShedLimit(handle http.HandlerFunc, maxConcurrent int) http.HandlerFunc {
	reqs := make(chan struct{}, maxConcurrent)
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case reqs <- struct{}{}:
			handle(w, r)
			<-reqs
		default:
			http.Error(w, "overloaded", 503)
		}
	}
}

func ShedLimitTimeout(handle http.HandlerFunc, maxConcurrent int, timeout time.Duration) http.HandlerFunc {
	reqs := make(chan struct{}, maxConcurrent)
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case reqs <- struct{}{}:
			handle(w, r)
			<-reqs
		case <-time.After(timeout):
			http.Error(w, "overloaded", 503)
		}
	}
}

func ShedLimitStack(handle http.HandlerFunc, maxConcurrent int, timeout time.Duration) http.HandlerFunc {
	reqs := make(chan struct{}, maxConcurrent)
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case reqs <- struct{}{}:
			handle(w, r)
			<-reqs
		case <-time.After(timeout):
			http.Error(w, "overloaded", 503)
		}
	}
}
