package user

import (
	"context"
	"net/http"
	"time"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	"github.com/ZangarZaynesh/forum/internal/domain"

	"github.com/ZangarZaynesh/forum/internal/module"
)

type handler struct {
	service domain.User
	ctx     context.Context
}

func NewHandler(ctx context.Context, user domain.User) handlers.User {
	return &handler{
		service: user,
		ctx:     ctx,
	}
}

func (h *handler) Register(router *http.ServeMux) {
	router.HandleFunc("/registration/", h.Registration)
	router.HandleFunc("/registration/created/", h.CreatedUser)
	router.HandleFunc("/auth/", h.SignIn)
	router.HandleFunc("/auth/success/", h.SignAccess)
}

func (h *handler) CreatedUser(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/registration/created/", "POST", w, r) {
		return
	}

	dto := new(module.CreateUserDTO)
	dto.Add(r)

	if !h.CheckLogin(dto, w, r) {
		return
	}

	if !h.CheckEmail(dto, w, r) {
		return
	}

	if !h.CheckPassword(dto, w, r) {
		return
	}

	if !dto.GeneratePassword() {
		// h.Error = http.StatusText(500)
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		// h.Error = ""
		return
	}

	if err := h.service.Create(h.ctx, dto); err != nil {
		// h.Error = http.StatusText(500)
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		// h.Error = ""
		return
	}
	// h.Error = "Successful"
	handlers.ExecTemp("templates/created.html", "created.html", w, r)
}

func (h *handler) Registration(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/registration/", "GET", w, r) {
		return
	}
	handlers.ExecTemp("templates/registration.html", "registration.html", w, r)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/auth/", "GET", w, r) {
		return
	}

	handlers.ExecTemp("templates/signIn.html", "signIn.html", w, r)
}

func (h *handler) SignAccess(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/auth/success/", "POST", w, r) {
		return
	}

	dto := new(module.SignUserDTO)
	dto.Add(r)

	if !h.CheckSignIn(dto, w, r) {
		return
	}

	dto.UUID, dto.CreateTimeUUID, dto.Duration = h.service.CreateCookie(w), time.Now(), time.Now().AddDate(0, 0, 1)
	if err := h.service.AddCookie(h.ctx, dto); err != nil {
		// h.Error = err.Error()
		handlers.ExecTemp("templates/error.html", "error.html", w, r)
		// h.Error = ""
		return
	}
	// handlers.ExecTemp("templates/index.html", "index.html", w, r)
	http.Redirect(w, r, "/", 302)
}
