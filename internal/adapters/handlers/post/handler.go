package post

import (
	"context"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	"github.com/ZangarZaynesh/forum/internal/domain"
	"github.com/ZangarZaynesh/forum/internal/module"
)

type handler struct {
	service domain.Post
	ctx     context.Context
	Error   string
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
		dto.UserId = 0
	}

	if err := h.service.GetPost(h.ctx, dto); err != nil {
		handlers.ExecTemp(err.Error(), "error.html", w, r)
		return
	}

	handlers.ExecTemp(dto, "index.html", w, r)
}

func (h *handler) CheckPathMethod(Path, Method string, w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != Path {
		handlers.ExecTemp(http.StatusText(400), "error.html", w, r)
		return false
	}

	if r.Method != Method {
		handlers.ExecTemp(http.StatusText(405), "error.html", w, r)
		return false
	}
	return true
}
