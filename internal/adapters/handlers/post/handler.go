package post

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	"github.com/ZangarZaynesh/forum/internal/domain"
	"github.com/ZangarZaynesh/forum/internal/module"
)

type handler struct {
	service domain.Post
	ctx     context.Context
}

func NewHandler(ctx context.Context, post domain.Post) handlers.Post {
	return &handler{
		service: post,
		ctx:     ctx,
	}
}

func (h *handler) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.Home)
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/", "GET", w, r) {
		return
	}
	dto := new(module.HomePageDTO)
	if err := h.CheckCookie(h.ctx, r, dto); err != nil {

		// h.Error = err.Error()
		if !errors.Is(err, sql.ErrNoRows) {
			handlers.ExecTemp("templates/error.html", "error.html", w, r)
			return
		}
		// napisat' owibku
		// h.Error = ""
	}

	if err := h.service.GetPost(h.ctx, dto); err != nil {
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		return
	}

}

func (h *handler) CheckPathMethod(Path, Method string, w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != Path {
		// h.Error = http.StatusText(400)
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		// h.Error = ""
		return false
	}

	if r.Method != Method {
		// h.Error = http.StatusText(405)
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		// h.Error = ""
		return false
	}
	return true
}
