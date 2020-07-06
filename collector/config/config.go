package config

import (
	"time"

	"github.com/spf13/viper"
)

func Init() {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_READ_TIMEOUT", time.Minute)
	viper.SetDefault("SERVER_WRITE_TIMEOUT", time.Minute)
	viper.SetDefault("SERVER_IDLE_TIMEOUT", time.Minute*10) // nolint:gomnd,gocritic

	viper.SetDefault("CLICKHOUSE_READ_TIMEOUT", 10)
	viper.SetDefault("CLICKHOUSE_WRITE_TIMEOUT", 20)
	viper.SetDefault("ENTRY_REPOSITORY_CAP", 1000)
	viper.SetDefault("ENTRY_REPOSITORY_PERIOD", time.Second)
}
