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

func (h *handler) CreatedUser(w http.ResponseWriter, r *http.Request) {

	if !CheckPathMethod("/registration/created/", "POST", w, r) {
		return
	}

	dto := new(module.CreateUserDTO)
	dto.Add(r)

	if !CheckLogin(h, dto, w, r) {
		return
	}

	if !CheckEmail(h, dto, w, r) {
		return
	}

	if !CheckPassword(h, dto, w, r) {
		return
	}

	if !dto.GeneratePassword() {
		handlers.ExecTemp("templates/error.html", "error.html", "500 Internal Server Error", w, r)
		return
	}

	if err := h.service.Create(h.ctx, dto); err != nil {
		handlers.ExecTemp("templates/error.html", "error.html", "500 Internal Server Error", w, r)
		return
	}

	handlers.ExecTemp("templates/created.html", "created.html", "Successful", w, r)
}

func (h *handler) Registration(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod("/registration/", "GET", w, r) {
		return
	}
	handlers.ExecTemp("templates/registration.html", "registration.html", "", w, r)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod("/auth/", "GET", w, r) {
		return
	}

	handlers.ExecTemp("templates/signIn.html", "signIn.html", "", w, r)
}

func (h *handler) SignAccess(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod("/registration/created/", "POST", w, r) {
		return
	}

	dto := new(module.SignUserDTO)
	dto.Add(r)

	if !CheckSignIn(h, dto, w, r) {
		return
	}

}
