package post

import (
	"context"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	"github.com/ZangarZaynesh/forum/internal/domain"
)

type handler struct {
	service domain.Post
	ctx     context.Context
	Error   string
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/", "GET", w, r) {
		return
	}

	if err := h.CheckCookie(r); err != nil {
		h.Error = err.Error()
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		h.Error = ""
		return
	}
}

func (h *handler) CheckPathMethod(Path, Method string, w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != Path {
		h.Error = http.StatusText(400)
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		h.Error = ""
		return false
	}

	if r.Method != Method {
		h.Error = http.StatusText(405)
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		h.Error = ""
		return false
	}
	return true
}
