package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gadavy/tracing"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/lissteron/loghole/dashboard/config"
	"github.com/lissteron/loghole/dashboard/internal/app/http/handlers"
	"github.com/lissteron/loghole/dashboard/internal/app/repositories/clickhouse"
	"github.com/lissteron/loghole/dashboard/internal/app/usecases"
	"github.com/lissteron/loghole/dashboard/pkg/clickhouseclient"
	"github.com/lissteron/loghole/dashboard/pkg/log"
	"github.com/lissteron/loghole/dashboard/pkg/server"
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

	logger.Info("initClickhouse start")

	// Init clients
	clickhousedb, err := initClickhouse()
	if err != nil {
		logger.Fatalf("init clickhouse db client failed: %v", err)
	}

	logger.Info("clickhouse inited")

	// Init repositories
	entryRepository := clickhouse.NewRepository(clickhousedb.Client(), traceLogger)

	// Init use cases
	entryList := usecases.NewListEntry(entryRepository, traceLogger)

	// Init http handlers
	var (
		listEntryHandlers = handlers.NewEntryHandlers(entryList, traceLogger)
		tracingMiddleware = handlers.NewTracingMiddleware(tracer)
	)

	// Init http server
	srv := initHTTPServer()

	// Init v1 routes
	r := srv.Router()
	r1 := r.PathPrefix("/api/v1").Subrouter()
	r1.Use(tracingMiddleware.Middleware)
	r1.HandleFunc("/entry/list", listEntryHandlers.ListEntryHandler).Methods("POST")

	var errGroup errgroup.Group

	errGroup.Go(func() error {
		logger.Infof("start http server on: %s", srv.Addr())
		return srv.ListenAndServe()
	})

	<-exit

	logger.Info("stopping application")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error while stopping web server: %v", err)
	}

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
		ServiceName: "dashboard_backend",
	})
}

func initClickhouse() (*clickhouseclient.Client, error) {
	return clickhouseclient.NewClient(
		viper.GetString("CLICKHOUSE_URI"),
		viper.GetString("CLICKHOUSE_USER"),
		viper.GetString("CLICKHOUSE_DATABASE"),
	)
}

func initHTTPServer() *server.HTTP {
	return server.NewHTTP(
		fmt.Sprintf("0.0.0.0:%s", viper.GetString("SERVICE_HTTP_PORT")),
		server.WithReadTimeout(time.Minute),
		server.WithWriteTimeout(time.Minute),
		server.WithIdleTimeout(time.Minute*10), // nolint:gomnd,gocritic
	)
}
