package middleware

import (
	"math/rand"
	"net/http"
	"time"
)

func RandomDelayMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		delay := rand.Intn(2000)
		time.Sleep(time.Duration(delay) * time.Millisecond)

		next.ServeHTTP(w, r)
	})
}
