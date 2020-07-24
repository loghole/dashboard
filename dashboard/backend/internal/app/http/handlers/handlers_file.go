// +build !bindata

package handlers

import (
	"net/http"
)

type FilesHandler struct {
	path string
}

func NewFilesHandler(path string) *FilesHandler {
	return &FilesHandler{path: path}
}

func (h *FilesHandler) Handler() http.Handler {
	return http.FileServer(http.Dir(h.path))
}
