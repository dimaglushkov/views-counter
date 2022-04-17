package app

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
)

const ext = ".svg"

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

	if !strings.HasSuffix(r.URL.Path, ext) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fullUrl := r.URL.Path[1:]
	urlHost := strings.Split(fullUrl, ".")[0]
	urlElems := strings.Split(fullUrl, "/")
	url := strings.Join(urlElems[:len(urlElems)-1], "/")
	badgeLabel := strings.TrimSuffix(urlElems[len(urlElems)-1], ext)
	if !strings.Contains(r.Referer(), urlHost) && !strings.Contains(r.UserAgent(), "github-camo") {
		w.WriteHeader(http.StatusForbidden)
		return
	}

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

	buf := buildBadge(badgeLabel, count)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
	w.Header().Set("ETag", fmt.Sprintf("%x", md5.Sum(buf.Bytes())))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf.Bytes())
}
