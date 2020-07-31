package handlers

import (
	"net"
	"net/http"
)

type RemoteIPMiddleware struct {
	header string
}

func NewRemoteIPMiddleware(header string) *RemoteIPMiddleware {
	return &RemoteIPMiddleware{header: header}
}

func (m *RemoteIPMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var remoteAddr string

		switch {
		case m.header != "":
			remoteAddr = r.Header.Get(m.header)
		default:
			remoteAddr = r.RemoteAddr
		}

		host, _, err := net.SplitHostPort(remoteAddr)
		if err == nil {
			remoteAddr = host
		}

		r.RemoteAddr = remoteAddr

		r.Header.Del(m.header)

		next.ServeHTTP(w, r)
	})
}
