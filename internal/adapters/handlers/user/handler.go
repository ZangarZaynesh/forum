package user

import (
	"context"
	"net/http"
	"time"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"

	"github.com/ZangarZaynesh/forum/internal/module"

	"github.com/ZangarZaynesh/forum/internal/domain"
)

type handler struct {
	service domain.User
	ctx     context.Context
	Error   string
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod(h, "/", "GET", w, r) {
		return
	}

	Session := r.Cookies()
}

func (h *handler) CreatedUser(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod(h, "/registration/created/", "POST", w, r) {
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
		h.Error = "500 Internal Server Error"
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		h.Error = ""
		return
	}

	if err := h.service.Create(h.ctx, dto); err != nil {
		h.Error = "500 Internal Server Error"
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		h.Error = ""
		return
	}
	h.Error = "Successful"
	handlers.ExecTemp("templates/created.html", "created.html", w, r)
}

func (h *handler) Registration(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod(h, "/registration/", "GET", w, r) {
		return
	}
	handlers.ExecTemp("templates/registration.html", "registration.html", w, r)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod(h, "/auth/", "GET", w, r) {
		return
	}

	handlers.ExecTemp("templates/signIn.html", "signIn.html", w, r)
}

func (h *handler) SignAccess(w http.ResponseWriter, r *http.Request) {
	if !CheckPathMethod(h, "/registration/created/", "POST", w, r) {
		return
	}

	dto := new(module.SignUserDTO)
	dto.Add(r)

	if !CheckSignIn(h, dto, w, r) {
		return
	}

	dto.UUID, dto.CreateTimeUUID, dto.Duration = h.service.CreateCookie(w), time.Now(), time.Now().AddDate(0, 0, 1)
	if err := h.service.AddCookie(h.ctx, dto); err != nil {
		h.Error = err.Error()
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		h.Error = ""
		return
	}
	handlers.ExecTemp("templates/index.html", "index.html", w, r)
}
