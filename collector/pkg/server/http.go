package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Option func(h *HTTP)

func WithReadTimeout(timeout time.Duration) Option {
	return func(h *HTTP) {
		h.server.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(h *HTTP) {
		h.server.WriteTimeout = timeout
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(h *HTTP) {
		h.server.IdleTimeout = timeout
	}
}

func WithTLS(certFile, keyFile string) Option {
	return func(h *HTTP) {
		h.cert = certFile
		h.key = keyFile
	}
}

type HTTP struct {
	cert   string
	key    string
	addr   string
	server *http.Server
	router *mux.Router
}

func NewHTTP(addr string, options ...Option) *HTTP {
	router := mux.NewRouter()

	server := &HTTP{
		addr:   addr,
		router: router,
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func (h *HTTP) ListenAndServe() (err error) {
	switch h.key != "" && h.cert != "" {
	case true:
		err = h.server.ListenAndServeTLS(h.cert, h.key)
	default:
		err = h.server.ListenAndServe()
	}

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (h *HTTP) Router() *mux.Router {
	return h.router
}

func (h *HTTP) Addr() string {
	return fmt.Sprintf("http://%s", h.addr)
}

func (h *HTTP) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
