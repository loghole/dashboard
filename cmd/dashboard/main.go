package main

import (
	"compress/gzip"
	log "log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	tron "github.com/loghole/tron"

	"github.com/loghole/dashboard/config"
	entryV1 "github.com/loghole/dashboard/internal/app/controllers/entry/v1"
	suggestV1 "github.com/loghole/dashboard/internal/app/controllers/suggest/v1"
	"github.com/loghole/dashboard/internal/app/controllers/ui"
	"github.com/loghole/dashboard/internal/app/repositories/clickhouse"
	"github.com/loghole/dashboard/internal/app/usecases"
	"github.com/loghole/dashboard/internal/pkg/clickhouseclient"
)

func main() {
	app, err := tron.New(tron.AddLogCaller(), tron.WithRealtimeConfig())
	if err != nil {
		log.Fatalf("can't create app: %s", err)
	}

	defer app.Close()

	// Init clients.
	clickhouseDB, err := clickhouseclient.NewClient(config.ClickhouseConfig())
	if err != nil {
		app.Logger().Fatalf("init clickhouse db client failed: %v", err)
	}

	// Init repositories.
	repository := clickhouse.NewRepository(clickhouseDB.Client(), app.TraceLogger())

	// Init use cases.
	var (
		entryList   = usecases.NewListEntry(repository, app.TraceLogger())
		suggestList = usecases.NewListSuggest(repository, app.TraceLogger())
	)

	var (
		entryV1Handler   = entryV1.NewImplementation(entryList)
		suggestV1Handler = suggestV1.NewImplementation(suggestList)
	)

	app.Router().Use(middleware.Compress(gzip.DefaultCompression))

	app.Router().HandleFunc("/", ui.NewRedirectHandler("/ui/"))
	app.Router().Mount("/ui", http.StripPrefix("/ui/", ui.NewHandler(config.GetFrontendPath())))

	if err := app.WithRunOptions(config.TLSKeyPair()).Run(entryV1Handler, suggestV1Handler); err != nil {
		app.Logger().Fatalf("can't run app: %v", err)
	}

	if err := clickhouseDB.Close(); err != nil {
		app.Logger().Errorf("close clickhouse failed: %v", err)
	}
}
