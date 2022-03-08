package post

import (
	"context"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/domain"
	"github.com/ZangarZaynesh/forum/internal/module"
)

type service struct {
	Post repository.Post
}

func NewService(post repository.Post) domain.Post {
	return &service{Post: post}
}

func (p *service) CheckCookie(ctx context.Context, session *http.Cookie, dto *module.HomePageDTO) error {
	if err := p.Post.CheckCookie(ctx, session, dto); err != nil {
		return err
	}
	return nil
}

func (p *service) GetPost(ctx context.Context, dto *module.HomePageDTO) error {
	if err := p.Post.GetPost(ctx, dto); err != nil {
		return err
	}
	return nil
}
