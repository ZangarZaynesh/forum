package repository

import (
	"context"

	"github.com/ZangarZaynesh/forum/internal/module"
)

type User interface {
	CheckByLogin(ctx context.Context, dto module.CreateUserDTO) error
	CheckByEmail(ctx context.Context, dto module.CreateUserDTO) error
	Create(ctx context.Context, dto *module.CreateUserDTO) error
	Delete(id int) error
}
