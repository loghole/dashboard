package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gadavy/tracing"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/lissteron/loghole/collector/config"
	"github.com/lissteron/loghole/collector/internal/app/controllers/http/handlers"
	"github.com/lissteron/loghole/collector/internal/app/repositories/clickhouse"
	"github.com/lissteron/loghole/collector/internal/app/services/entry"
	"github.com/lissteron/loghole/collector/pkg/clickhouseclient"
	"github.com/lissteron/loghole/collector/pkg/log"
	"github.com/lissteron/loghole/collector/pkg/server"
)

// nolint: funlen,gocritic
func main() {
	// Init config, logger, exit chan
	config.Init()

	logger, err := initLogger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "init logger failed: %v", err)
		os.Exit(1)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	// Init jaeger tracer.
	tracer, err := initTracer()
	if err != nil {
		logger.Fatalf("init tracing client failed: %v", err)
	}

	traceLogger := tracing.NewTraceLogger(logger)

	// Init clients
	clickhousedb, err := initClickhouse()
	if err != nil {
		logger.Fatalf("init clickhouse db client failed: %v", err)
	}

	// Init repository
	repository := clickhouse.NewEntryRepository(
		clickhousedb.Client(),
		traceLogger,
		viper.GetInt("ENTRY_REPOSITORY_CAP"),
		viper.GetDuration("ENTRY_REPOSITORY_PERIOD"),
	)

	// Init service
	entryService := entry.NewService(repository, traceLogger)

	// Init handlers
	entryHandlers := handlers.NewEntryHandlers(entryService, traceLogger, tracer)

	srv := initHTTPServer()

	r := srv.Router()
	r.HandleFunc("/api/v1/store", entryHandlers.StoreItemHandler)
	r.HandleFunc("/api/v1/store/list", entryHandlers.StoreListHandler)
	r.HandleFunc("/api/v1/ping", entryHandlers.PingHandler)

	var errGroup, ctx = errgroup.WithContext(context.Background())

	errGroup.Go(func() error {
		logger.Info("start entry writer")
		return repository.Run(ctx)
	})

	errGroup.Go(func() error {
		logger.Infof("start http server on: %s", srv.Addr())
		return srv.ListenAndServe()
	})

	select {
	case <-exit:
		logger.Info("stopping application")
	case <-ctx.Done():
		logger.Error("stopping application with error")
	}

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error while stopping web server: %v", err)
	}

	repository.Stop()

	if err = errGroup.Wait(); err != nil {
		logger.Errorf("error while waiting for goroutines: %v", err)
	}

	if err = tracer.Close(); err != nil {
		logger.Errorf("error while stopping tracer: %v", err)
	}

	if err = clickhousedb.Close(); err != nil {
		logger.Errorf("error while stopping clickhouse db: %v", err)
	}

	logger.Info("application stopped")
}

func initLogger() (*zap.SugaredLogger, error) {
	return log.NewLogger(
		log.SetLevel(viper.GetString("LOGGER_LEVEL")),
		log.AddCaller(),
	)
}

func initTracer() (*tracing.Tracer, error) {
	return tracing.NewTracer(&tracing.Config{
		URI:         viper.GetString("JAEGER_URI"),
		Enabled:     viper.GetString("JAEGER_URI") != "",
		ServiceName: "collector",
	})
}

func initClickhouse() (*clickhouseclient.Client, error) {
	return clickhouseclient.NewClient(&clickhouseclient.Options{
		Addr:         viper.GetString("CLICKHOUSE_URI"),
		User:         viper.GetString("CLICKHOUSE_USER"),
		Database:     viper.GetString("CLICKHOUSE_DATABASE"),
		ReadTimeout:  viper.GetInt("CLICKHOUSE_READ_TIMEOUT"),
		WriteTimeout: viper.GetInt("CLICKHOUSE_WRITE_TIMEOUT"),
		SchemaPath:   viper.GetString("CLICKHOUSE_SCHEMA_PATH"),
	})
}

func initHTTPServer() *server.HTTP {
	return server.NewHTTP(
		fmt.Sprintf("0.0.0.0:%s", viper.GetString("SERVER_HTTP_PORT")),
		server.WithReadTimeout(viper.GetDuration("SERVER_READ_TIMEOUT")),
		server.WithWriteTimeout(viper.GetDuration("SERVER_WRITE_TIMEOUT")),
		server.WithIdleTimeout(viper.GetDuration("SERVER_IDLE_TIMEOUT")),
		server.WithTLS(viper.GetString("SERVER_CERT"), viper.GetString("SERVER_KEY")),
	)
}
