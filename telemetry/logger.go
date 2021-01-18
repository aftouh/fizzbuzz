package telemetry

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Level       string
	Environment string
	ServiceName string
}

// InitLogger creates a new  logger and overrides the global zap logger with the created one
// The logger is accessible with zap.L()
func InitLogger(c LoggerConfig) {
	level, err := getZapLogLevel(c.Level)
	if err != nil {
		panic(fmt.Errorf("log level error: %s", err))
	}

	// Logger encoder
	var encoder zapcore.Encoder
	if c.Environment == "local" {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, os.Stdout, level)

	logger := zap.New(core, zap.AddCaller(),
		zap.Fields(
			zap.String("service", c.ServiceName),
			zap.String("env", c.Environment),
		),
	)

	zap.ReplaceGlobals(logger)
}

func LoggerChiMiddleware(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.Info("Served",
					zap.String("proto", r.Proto),
					zap.String("path", r.URL.Path),
					zap.Duration("lat", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
					zap.String("reqId", middleware.GetReqID(r.Context())))
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

func getZapLogLevel(s string) (l zapcore.LevelEnabler, err error) {
	switch s {
	case "debug", "DEBUG":
		l = zapcore.DebugLevel
	case "info", "INFO":
		l = zapcore.InfoLevel
	case "warn", "WARN":
		l = zapcore.WarnLevel
	case "error", "ERROR", "": // zero value sets log level to error
		l = zapcore.ErrorLevel
	case "fatal", "FATAL":
		l = zapcore.FatalLevel
	default:
		return nil, fmt.Errorf("unknown log level %q", s)
	}
	return
}
