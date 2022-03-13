package user

import (
	"context"
	"net/http"
	"time"

	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/domain"
	"github.com/ZangarZaynesh/forum/internal/module"
	"github.com/satori/uuid"
)

type service struct {
	user repository.User
}

func NewService(user repository.User) domain.User {
	return &service{user: user}
}

// Create User
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

// Cookie
func (u *service) CreateCookie(w http.ResponseWriter) uuid.UUID {
	sessionID := uuid.NewV4()
	cookie := &http.Cookie{
		Name:    "session",
		Path:    "/",
		Value:   sessionID.String(),
		Expires: time.Now().AddDate(0, 0, 1),
		MaxAge:  24 * 3600,
	}
	http.SetCookie(w, cookie)
	return sessionID
}

func (u *service) AddCookie(ctx context.Context, dto *module.SignUserDTO) error {
	if err := u.user.AddCookie(ctx, dto); err != nil {
		return err
	}
	return nil
}

func (u *service) DeleteCookie(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	session := http.Cookie{}
	session.Name = "session"
	session.MaxAge = -1
	session.Expires = time.Unix(1, 0)
	session.Path = "/"
	http.SetCookie(w, &session)
}

func (u *service) DeleteUUID(ctx context.Context, dto *module.HomePageDTO) error {
	if err := u.user.DeleteUUID(ctx, dto); err != nil {
		return err
	}
	return nil
}

func (u *service) CheckCookie(ctx context.Context, session *http.Cookie, dto *module.HomePageDTO) error {
	if err := u.user.CheckCookie(ctx, session, dto); err != nil {
		return err
	}
	return nil
}
