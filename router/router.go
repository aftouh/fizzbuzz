package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httptracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/bridge/opentracing"
	"go.uber.org/zap"

	"github.com/aftouh/fizzbuzz/config"
	"github.com/aftouh/fizzbuzz/handlers"
	"github.com/aftouh/fizzbuzz/telemetry"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	// Injects a request ID into the context of each request
	r.Use(middleware.RequestID)
	// Sets a http.Request's RemoteAddr to either X-Forwarded-For or X-Real-IP
	r.Use(middleware.RealIP)
	// Gracefully absorb panics and prints the stack trace
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
		},
	}))
	r.Use(telemetry.LoggerChiMiddleware(zap.L()))

	cfg := config.GetConfig()
	if cfg.Tracer.Enabled {
		// Convert the opentracing tracer to an opentelemetry tracer
		tracer := opentracing.NewBridgeTracer()
		tracer.SetOpenTelemetryTracer(otel.GetTracerProvider().Tracer("bridge"))
		r.Use(httptracer.Tracer(tracer, httptracer.Config{
			ServiceName: cfg.ServiceName,
			SkipFunc: func(r *http.Request) bool {
				return r.URL.Path == "/healthz" || r.URL.Path == "/metrics"
			},
		}))
	}

	r.Route("/v1", func(r chi.Router) {
		r.Get("/fizzbuzz", handlers.Fizzbuzz)
	})

	return r
}
