package user

import (
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	"github.com/ZangarZaynesh/forum/internal/module"
)

func (h *handler) CheckLogin(dto *module.CreateUserDTO, w http.ResponseWriter, r *http.Request) bool {
	if dto.Login == "" {
		h.Error = "Enter login"
		http.Redirect(w, r, "/registration/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}

	if err := h.service.CheckByLogin(h.ctx, dto); err != nil {
		h.Error = err.Error()
		http.Redirect(w, r, "/registration/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}
	return true
}

func (h *handler) CheckEmail(dto *module.CreateUserDTO, w http.ResponseWriter, r *http.Request) bool {
	if !dto.IsEmailValid() {
		h.Error = "Incorrected email"
		http.Redirect(w, r, "/registration/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}

	if err := h.service.CheckByEmail(h.ctx, dto); err != nil {
		h.Error = err.Error()
		http.Redirect(w, r, "/registration/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}
	return true
}

func (h *handler) CheckPassword(dto *module.CreateUserDTO, w http.ResponseWriter, r *http.Request) bool {
	if dto.Password == "" {
		h.Error = "Enter password"
		http.Redirect(w, r, "/registration/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}

	if !dto.CheckPassConfirm() {
		h.Error = "Incorrected confirm"
		http.Redirect(w, r, "/registration/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}
	return true
}

func (h *handler) CheckSignIn(dto *module.SignUserDTO, w http.ResponseWriter, r *http.Request) bool {
	if dto.Login == "" {
		h.Error = "Enter login"
		http.Redirect(w, r, "/auth/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}

	if err := h.service.CheckSignIn(h.ctx, dto); err != nil {
		h.Error = err.Error()
		http.Redirect(w, r, "/auth/", http.StatusPermanentRedirect)
		h.Error = ""
		return false
	}
	return true
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
