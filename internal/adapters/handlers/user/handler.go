package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	"github.com/ZangarZaynesh/forum/internal/domain"

	"github.com/ZangarZaynesh/forum/internal/module"
)

type handler struct {
	service  domain.User
	ctx      context.Context
	Error    string
	UserAuth string
}

func NewHandler(ctx context.Context, user domain.User) handlers.User {
	return &handler{
		service: user,
		ctx:     ctx,
	}
}

func (h *handler) Register(router *http.ServeMux) {
	fmt.Println("/registration/")
	router.HandleFunc("/registration/", h.Registration)
	fmt.Println("/registration/created/")
	router.HandleFunc("/registration/created/", h.CreatedUser)
	fmt.Println("/auth/")
	router.HandleFunc("/auth/", h.SignIn)
	fmt.Println("/auth/user/")
	router.HandleFunc("/auth/user/", h.SignAccess)
	fmt.Println("/signOut/")
	router.HandleFunc("/signOut/", h.SignOut)
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
		handlers.ExecTemp(http.StatusText(500), "error.html", w, r)
		return
	}

	if err := h.service.Create(h.ctx, dto); err != nil {
		handlers.ExecTemp(http.StatusText(500), "error.html", w, r)
		return
	}

	handlers.ExecTemp("Successful", "created.html", w, r)
}

func (h *handler) Registration(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/registration/", "GET", w, r) {
		return
	}
	handlers.ExecTemp(nil, "registration.html", w, r)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/auth/", "GET", w, r) {
		return
	}

	handlers.ExecTemp(nil, "signIn.html", w, r)
}

func (h *handler) SignAccess(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/auth/user/", "POST", w, r) {
		return
	}

	dto := new(module.SignUserDTO)
	dto.Add(r)

	if !h.CheckSignIn(dto, w, r) {
		return
	}

	h.service.DeleteCookie(h.ctx, w)

	if err := h.service.DeleteSession(h.ctx, dto.UserId); err != nil {
		handlers.ExecTemp(err.Error(), "error.html", w, r)
		return
	}

	dto.UUID, dto.CreateTimeUUID, dto.Duration = h.service.CreateCookie(w), time.Now(), time.Now().AddDate(0, 0, 1)
	if err := h.service.AddCookie(h.ctx, dto); err != nil {
		handlers.ExecTemp(err.Error(), "error.html", w, r)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func (h *handler) SignOut(w http.ResponseWriter, r *http.Request) {
	if !h.CheckPathMethod("/signOut/", "GET", w, r) {
		return
	}

	dto := new(module.HomePageDTO)
	if err := h.CheckCookie(h.ctx, r, dto); err != nil {
		http.Redirect(w, r, "/", 308)
		return
	}

	if err := h.DeleteCookie(h.ctx, w, r, dto); err != nil {
		handlers.ExecTemp(err.Error(), "error.html", w, r)
		return
	}
	http.Redirect(w, r, "/", 308)
}
