package main

import (
	"net/http"
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
