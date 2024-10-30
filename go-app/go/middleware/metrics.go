package middleware

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Общее число запросов",
		},
		[]string{"method", "endpoint"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Длительность запросов в секундах",
			Buckets: []float64{0.05, 0.1, 0.5, 1, 1.5, 2},
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	// Регистрация метрик
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		endpoint := r.URL.Path

		start := time.Now()

		delay := rand.Intn(1500)
		time.Sleep(time.Duration(delay) * time.Millisecond)

		next.ServeHTTP(w, r)

		if r.URL.Path == "/metrics" {
			return
		}
		requestCounter.WithLabelValues(method, endpoint).Inc()
		requestDuration.WithLabelValues(method, endpoint).Observe(time.Since(start).Seconds())
	})
}
