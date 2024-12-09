package middleware

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
)

func TracesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			next.ServeHTTP(w, r)
			return
		}

		tracer := otel.Tracer("http-tracer")
		ctx, span := tracer.Start(r.Context(), r.Method+" "+r.URL.Path)
		defer span.End()

		_, span1 := tracer.Start(ctx, "pre-action")
		delay := rand.Intn(400)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		span1.End()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			_, span2 := tracer.Start(ctx, "background_work")
			delay := rand.Intn(400)
			time.Sleep(time.Duration(delay) * time.Millisecond)
			span2.End()
		}()

		go func() {
			defer wg.Done()

			_, span3 := tracer.Start(ctx, "serving_request")
			next.ServeHTTP(w, r.WithContext(ctx))
			span3.End()
		}()

		wg.Wait()
		span.End()
	})
}
