package config

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/loghole/lhw/zap"
	"github.com/loghole/tracing"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"

	"github.com/loghole/dashboard/pkg/clickhouseclient"
	"github.com/loghole/dashboard/pkg/server"
)

const (
	defaultServiceName = "dashboard"
)

// nolint:gochecknoglobals // build args
var (
	InstanceUUID = uuid.New()
	ServiceName  string
	AppName      string
	GitHash      string
	Version      string
	BuildAt      string
)

func Init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("json")
	viper.SetConfigName(defaultServiceName)
	viper.AddConfigPath("./configs/")

	_ = viper.ReadInConfig()

	viper.SetDefault("clickhouse.read.timeout", 10)
	viper.SetDefault("clickhouse.write.timeout", 20)
	viper.SetDefault("frontend.path", "")
}

func ClickhouseConfig() *clickhouseclient.Config {
	return &clickhouseclient.Config{
		Addr:         viper.GetString("clickhouse.uri"),
		User:         viper.GetString("clickhouse.user"),
		Password:     viper.GetString("clickhouse.password"),
		Database:     viper.GetString("clickhouse.database"),
		ReadTimeout:  viper.GetInt("clickhouse.read.timeout"),
		WriteTimeout: viper.GetInt("clickhouse.write.timeout"),
	}
}

func TracerConfig() *config.Configuration {
	return tracing.DefaultConfiguration(serviceName(), viper.GetString("jaeger.uri"))
}

func ServerConfig() *server.Config {
	return &server.Config{
		Addr:         fmt.Sprintf("0.0.0.0:%s", viper.GetString("server.http.port")),
		ReadTimeout:  viper.GetDuration("server.read.timeout"),
		WriteTimeout: viper.GetDuration("server.write.timeout"),
		IdleTimeout:  viper.GetDuration("server.idle.timeout"),
		TLSCertFile:  viper.GetString("server.tls.cert"),
		TLSKeyFile:   viper.GetString("server.tls.key"),
	}
}

func LoggerConfig() *zap.Config {
	return &zap.Config{
		Level:         viper.GetString("logger.level"),
		CollectorURL:  viper.GetString("logger.collector.url"),
		Namespace:     viper.GetString("logger.namespace"),
		Source:        serviceName(),
		BuildCommit:   GitHash,
		DisableStdout: false,
	}
}

func serviceName() string {
	switch {
	case ServiceName != "":
		return ServiceName
	case viper.GetString("service.name") != "":
		return viper.GetString("service.name")
	default:
		return defaultServiceName
	}
}
