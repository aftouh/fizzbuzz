package telemetry

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aftouh/fizzbuzz/config"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/metric/prometheus"
	"go.uber.org/zap"
)

func InitMeter() {
	exporter, err := prometheus.InstallNewPipeline(prometheus.Config{})
	if err != nil {
		zap.L().Panic("Failed to initialize prometheus exporter", zap.Error(err))
	}
	http.HandleFunc("/metrics", exporter.ServeHTTP)
	cfg := config.GetConfig()
	go func() {
		_ = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Metrics.Port), nil)
	}()

	zap.L().Info(fmt.Sprintf("Start prometheus server on port %d", cfg.Metrics.Port))

	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		zap.L().Panic("Failed starting runtime metric collector", zap.Error(err))
	}
}
