package user

import (
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"

	"github.com/ZangarZaynesh/forum/internal/module"
)

func CheckLogin(h *handler, dto *module.CreateUserDTO, w http.ResponseWriter, r *http.Request) bool {
	if dto.Login == "" {
		h.Error = "Enter login"
		http.Redirect(w, r, "/registration/", 302)
		h.Error = ""
		return false
	}

	if err := h.service.CheckByLogin(h.ctx, dto); err != nil {
		h.Error = err.Error()
		http.Redirect(w, r, "/registration/", 302)
		h.Error = ""
		return false
	}
	return true
}

func CheckEmail(h *handler, dto *module.CreateUserDTO, w http.ResponseWriter, r *http.Request) bool {
	if !dto.IsEmailValid() {
		h.Error = "Incorrected email"
		http.Redirect(w, r, "/registration/", 302)
		h.Error = ""
		return false
	}

	if err := h.service.CheckByEmail(h.ctx, dto); err != nil {
		h.Error = err.Error()
		http.Redirect(w, r, "/registration/", 302)
		h.Error = ""
		return false
	}
	return true
}

func CheckPassword(h *handler, dto *module.CreateUserDTO, w http.ResponseWriter, r *http.Request) bool {
	if dto.Password == "" {
		h.Error = "Enter password"
		http.Redirect(w, r, "/registration/", 302)
		h.Error = ""
		return false
	}

	if !dto.CheckPassConfirm() {
		h.Error = "Incorrected confirm"
		http.Redirect(w, r, "/registration/", 302)
		h.Error = ""
		return false
	}
	return true
}

func CheckPathMethod(Path, Method string, w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != Path {
		handlers.ExecTemp("templates/error.html", "error.html", "400 Bad Request", w, r)
		return false
	}

	if r.Method != Method {
		handlers.ExecTemp("templates/error.html", "error.html", "405 Method Not Allowed", w, r)
		return false
	}
	return true
}
