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
	row := r.db.QueryRow("SELECT user_id FROM sessions where key = ? ;", session.Value)
	err := row.Scan(&dto.UserId)
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}
	return nil
}

func (r *repo) GetPost(ctx context.Context, dto *module.HomePageDTO) error {
	rows, err := r.db.Query("SELECT date, user_id, title, post FROM posts")
	if err != nil {
		return err
	}

	for rows.Next() {
		var userId int
		var date time.Time
		var title, post, login string
		var authPost rune
		err = rows.Scan(&date, &userId, &title, &post)
		if err != nil {
			return err
		}

		row := r.db.QueryRow("SELECT login FROM users where id = ?", userId)

		err = row.Scan(&login)
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}

		if dto.UserId == userId {
			authPost = 't'
		} else {
			authPost = 'f'
		}

		dto.Posts = append(dto.Posts, module.ShowPostDTO{
			Title:    title,
			Post:     post,
			Date:     date,
			Login:    login,
			AuthPost: authPost,
		})
	}
	return nil
}
