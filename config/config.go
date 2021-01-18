package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServiceName    string
	Environment    string
	Port           int
	LogLevel       string
	AllowedOrigins []string
	Metrics        MetricsConfig
	Tracer         TracerConfig
}
type MetricsConfig struct {
	Enabled bool
	Port    int
}
type TracerConfig struct {
	Enabled bool
}

var config AppConfig
var once sync.Once

const (
	configFilePathVarName = "CONFIG_FILE"
	defaultConfigFilePath = "config.yaml"
)

func GetConfig() AppConfig {
	once.Do(func() {
		v := viper.New()
		// Set default values
		v.SetDefault("serviceName", "fizzbuzz")
		v.SetDefault("port", 8080)
		v.SetDefault("logLevel", "error")
		v.SetDefault("allowedOrigins", []string{"*"})
		v.SetDefault("metrics", MetricsConfig{
			Enabled: true,
			Port:    8082,
		})
		v.SetDefault("tracer", TracerConfig{
			Enabled: true,
		})

		var filePath string
		if os.Getenv(configFilePathVarName) != "" {
			filePath = os.Getenv(configFilePathVarName)
		} else {
			filePath = defaultConfigFilePath
		}
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			panic(err)
		}

		// Read config file
		v.SetConfigFile(filePath)

		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}

		if err := v.UnmarshalExact(&config); err != nil {
			panic(fmt.Errorf("fatal error config parsing: %s", err))
		}
	})

	return config
}
