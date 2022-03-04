package post

import (
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/module"
)

func (h *handler) CheckCookie(r *http.Request, dto *module.HomePageDTO) error {
	session, err := r.Cookie("session")
	if err != nil {
		return err
	}
	
	if session.Value ==  {

	}
	return nil
}
