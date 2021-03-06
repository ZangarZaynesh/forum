package repository

import (
	"context"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/module"
)

type User interface {
	CheckByLogin(ctx context.Context, dto *module.CreateUserDTO) error
	CheckByEmail(ctx context.Context, dto *module.CreateUserDTO) error
	Create(ctx context.Context, dto *module.CreateUserDTO) error
	CheckSignIn(ctx context.Context, dto *module.SignUserDTO) error
	AddCookie(ctx context.Context, dto *module.SignUserDTO) error
	CheckCookie(ctx context.Context, session *http.Cookie, dto *module.HomePageDTO) error
	DeleteSession(ctx context.Context, userId int) error
}

type Post interface {
	CheckCookie(ctx context.Context, session *http.Cookie, dto *module.HomePageDTO) error
	GetPost(ctx context.Context, dto *module.HomePageDTO) error
}
