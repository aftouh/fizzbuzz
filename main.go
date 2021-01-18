package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/aftouh/fizzbuzz/config"
	"github.com/aftouh/fizzbuzz/router"
	"github.com/aftouh/fizzbuzz/telemetry"
)

func main() {
	cfg := config.GetConfig()

	telemetry.InitLogger(telemetry.LoggerConfig{
		Level:       cfg.LogLevel,
		Environment: cfg.Environment,
		ServiceName: cfg.ServiceName,
	})
	// flushing any buffered log entries before exiting
	defer zap.L().Sync()

	if cfg.Metrics.Enabled {
		telemetry.InitMeter()
	}
	if cfg.Tracer.Enabled {
		telemetry.InitTracer()
	}

	r := router.InitRouter()
	zap.L().Info(fmt.Sprintf("Start listing on port %d", cfg.Port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
		panic(fmt.Errorf("failed to start http server: %s", err))
	}
}
