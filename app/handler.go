package app

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
)

const urlSuffix = "/views.svg"

type handler struct {
	repo Repository
}

func NewHandler(repo Repository) *handler {
	return &handler{repo: repo}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		break
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasSuffix(r.URL.Path, urlSuffix) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url := strings.TrimSuffix(r.URL.Path[1:], urlSuffix)
	//urlHost := strings.Split(url, "/")[0]
	//if !strings.Contains(r.Referer(), urlHost) {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}

	count, err := h.repo.Visit(r.Context(), url)
	if err != nil {
		if _, ok := err.(*UnknownUrlError); ok {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	buf := buildBadge("views", count)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
	w.Header().Set("ETag", fmt.Sprintf("%x", md5.Sum(buf.Bytes())))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf.Bytes())
}
