package post

import (
	"context"
	"errors"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/module"
)

func (h *handler) CheckCookie(ctx context.Context, r *http.Request, dto *module.HomePageDTO) error {
	session, err := r.Cookie("session")
	if errors.Is(err, http.ErrNoCookie) {
		return err
	}

	if err = h.service.CheckCookie(ctx, session, dto); err != nil {
		return err
	}
	return nil
}
