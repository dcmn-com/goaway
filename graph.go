package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	promRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests",
			Help: "number of requests",
		},
	)
	promOpenRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "openrequests",
			Help: "number of open requests",
		},
	)
	promResptime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "responsetime",
			Help: "response time (ms)",
			Buckets: []float64{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				15, 20, 25, 30, 35, 40, 45, 50,
				60, 70, 80, 90, 100,
				150, 200, 250,
			},
		},
	)
)

func init() {
	prometheus.MustRegister(promRequests)
	prometheus.MustRegister(promOpenRequests)
	prometheus.MustRegister(promResptime)
	http.Handle("/metrics", prometheus.Handler())
}

func HandleFunc(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		promRequests.Inc()
		promOpenRequests.Inc()
		defer promOpenRequests.Dec()
		start := time.Now()
		defer promResptime.Observe(time.Since(start).Seconds() * 1000)
		handler(w, r)
	})
}
