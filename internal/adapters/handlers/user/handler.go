package user

import (
	"context"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"

	"github.com/ZangarZaynesh/forum/internal/module"

	"github.com/ZangarZaynesh/forum/internal/domain"
)

type handler struct {
	service domain.User
	ctx     context.Context
	Error   string
}

func (h *handler) Registration(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod("/registration/", "POST", w, r) {
		return
	}

	dto := new(module.CreateUserDTO)
	dto.Add(r)

	if !CheckLogin(*h, *dto, w, r) {
		return
	}

	if !CheckEmail(*h, *dto, w, r) {
		return
	}

	if !CheckPassword(*h, *dto, w, r) {
		return
	}

	if !dto.GeneratePassword() {
		handlers.ExecTemp("templates/error.html", "error.html", "500 Internal Server Error", w, r)
		return
	}

	h.service.Create(h.ctx, dto)
}
