package middleware

import (
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
			Name: "api_request_duration_seconds",
			Help: "Длительность запросов в секундах",
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

		next.ServeHTTP(w, r)

		requestCounter.WithLabelValues(method, endpoint).Inc()
		requestDuration.WithLabelValues(method, endpoint).Observe(time.Since(start).Seconds())
	})
}
