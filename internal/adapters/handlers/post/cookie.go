package post

import (
	"net/http"
)

func (h *handler) CheckCookie(r *http.Request) error {
	session, err := r.Cookie("session")
	if err != nil {
		return err
	}
	
	if session.Value ==  {

	}
	return nil
}
