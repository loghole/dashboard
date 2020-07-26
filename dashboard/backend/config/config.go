package config

import (
	"github.com/spf13/viper"
)

// nolint:gochecknoglobals // build args
var (
	ServiceName string
	AppName     string
	GitHash     string
	Version     string
	BuildAt     string
)

func Init() {
	viper.AutomaticEnv()

	viper.SetDefault("CLICKHOUSE_READ_TIMEOUT", 10)
	viper.SetDefault("CLICKHOUSE_WRITE_TIMEOUT", 20)
	viper.SetDefault("FRONTEND_PATH", "")
}
