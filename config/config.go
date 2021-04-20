// Code generated by tron. Place config helpers here.
package config

import (
	"github.com/loghole/tron"
	"github.com/spf13/viper"

	"github.com/loghole/dashboard/internal/pkg/clickhouseclient"
)

func ClickhouseConfig() *clickhouseclient.Config {
	return &clickhouseclient.Config{
		Addr:         viper.GetString(ClickhouseUri),
		User:         viper.GetString(ClickhouseUser),
		Password:     viper.GetString(ClickhousePassword),
		Database:     viper.GetString(ClickhouseDatabase),
		ReadTimeout:  viper.GetInt(ClickhouseReadTimeout),
		WriteTimeout: viper.GetInt(ClickhouseWriteTimeout),
	}
}

func TLSKeyPair() tron.RunOption {
	return tron.WithTLSKeyPair(
		viper.GetString(ServiceTlsKey),
		viper.GetString(ServiceTlsCert),
	)
}

func GetFrontendPath() string {
	return viper.GetString(FrontendPath)
}