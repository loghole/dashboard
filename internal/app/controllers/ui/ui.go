package ui

import (
	"net/http"
)

func NewHandler(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}

func NewRedirectHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = path

		http.Redirect(w, r, r.URL.String(), http.StatusPermanentRedirect)
	}
}
