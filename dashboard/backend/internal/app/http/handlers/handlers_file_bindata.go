// +build bindata

package handlers

import (
	"net/http"

	"github.com/lissteron/loghole/dashboard/bindata"
)

type FilesHandler struct{}

func NewFilesHanlder(_ string) *FilesHandler {
	return &FilesHandler{}
}

func (h *FilesHandler) Handler() http.Handler {
	return http.FileServer(bindata.AssetFile())
}
