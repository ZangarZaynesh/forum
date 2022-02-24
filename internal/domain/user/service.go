package user

import (
	"context"

	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/module"
)

type service struct {
	user repository.User
}

func (u *service) CheckByLogin(ctx context.Context, dto *module.CreateUserDTO) error {
	if err := u.user.CheckByLogin(ctx, dto); err != nil {
		return err
	}
	return nil
}

func (u *service) CheckSignIn(ctx context.Context, dto *module.SignUserDTO) error {
	if err := u.user.CheckSignIn(ctx, dto); err != nil {
		return err
	}
	return nil
}

func (u *service) CheckByEmail(ctx context.Context, dto *module.CreateUserDTO) error {
	if err := u.user.CheckByEmail(ctx, dto); err != nil {
		return err
	}
	return nil
}

func (u *service) Create(ctx context.Context, dto *module.CreateUserDTO) error {
	if err := u.user.Create(ctx, dto); err != nil {
		return err
	}
	return nil
}
