// +build !bindata

package handlers

import (
	"log"
	"net/http"
)

type FilesHandler struct {
	path string
}

func NewFilesHanlder(path string) *FilesHandler {
	log.Println("not bindata")

	return &FilesHandler{path: path}
}

func (h *FilesHandler) Handler() http.Handler {
	return http.FileServer(http.Dir(h.path))
}
