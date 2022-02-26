package domain

import (
	"context"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/module"
)

// interfaces from domain for api

type User interface {
	Create(ctx context.Context, dto *module.CreateUserDTO) error
	CheckByLogin(ctx context.Context, dto *module.CreateUserDTO) error
	CheckByEmail(ctx context.Context, dto *module.CreateUserDTO) error
	CheckSignIn(ctx context.Context, dto *module.SignUserDTO) error
	CreareCookie(w http.ResponseWriter)
}

type Post interface {
}

type Comment interface {
}

type Session interface {
}
