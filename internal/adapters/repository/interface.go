package repository

import (
	"context"

	"github.com/ZangarZaynesh/forum/internal/module"
)

type User interface {
	CheckByLogin(ctx context.Context, dto *module.CreateUserDTO) error
	CheckByEmail(ctx context.Context, dto *module.CreateUserDTO) error
	Create(ctx context.Context, dto *module.CreateUserDTO) error
	CheckSignIn(ctx context.Context, dto *module.SignUserDTO) error
	AddCookie(ctx context.Context, dto *module.SignUserDTO) error
}
