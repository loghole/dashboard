// +build bindata

package handlers

import (
	"net/http"

	"github.com/loghole/dashboard/bindata"
)

type FilesHandlers struct{}

func NewFilesHandlers(_ string) *FilesHandlers {
	return &FilesHandlers{}
}

func (h *FilesHandlers) Handler() http.Handler {
	return http.FileServer(bindata.AssetFile())
}
