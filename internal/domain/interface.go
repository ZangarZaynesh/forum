package domain

import (
	"context"
	"net/http"

	"github.com/satori/uuid"

	"github.com/ZangarZaynesh/forum/internal/module"
)

// interfaces from domain for api

type User interface {
	Create(ctx context.Context, dto *module.CreateUserDTO) error
	CheckByLogin(ctx context.Context, dto *module.CreateUserDTO) error
	CheckByEmail(ctx context.Context, dto *module.CreateUserDTO) error
	CheckSignIn(ctx context.Context, dto *module.SignUserDTO) error
	CreateCookie(w http.ResponseWriter) uuid.UUID
	AddCookie(ctx context.Context, dto *module.SignUserDTO) error
}

type Post interface{}

type Comment interface{}

type Session interface{}
