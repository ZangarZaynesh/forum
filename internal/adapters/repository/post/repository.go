package post

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

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

func (r *repo) GetPost(ctx context.Context, dto *module.HomePageDTO) error {
	rows, err := r.db.Query("SELECT id, date, user_id, title, post FROM posts")
	if err != nil {
		return err
	}

	for rows.Next() {
		var id, userId int
		var date time.Time
		var title, post string
		err = rows.Scan(&id, &date, &userId, &title, &post)
		if err != nil {
			return err
		}
		dto.Posts = append(dto.Posts, module.ShowPostDTO{
			Id:     id,
			Title:  title,
			Post:   post,
			Date:   date,
			UserId: userId,
		})
	}
	return nil
}
