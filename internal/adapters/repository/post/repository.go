package post

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/module"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository.Post {
	return &repo{db: db}
}

func (r *repo) CheckCookie(ctx context.Context, session *http.Cookie, dto *module.HomePageDTO) error {
	id := r.db.QueryRow("SELECT user_id FROM sessions where key = ? ;", session.Value)
	err := id.Scan(&dto.UserId)
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}
	return nil
}
