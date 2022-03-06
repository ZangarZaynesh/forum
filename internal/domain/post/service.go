package post

import (
	"context"
	"net/http"

	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/module"
)

type service struct {
	post repository.Post
}

func (p *service) CheckCookie(ctx context.Context, session *http.Cookie, dto *module.HomePageDTO) error {
	if err := p.post.CheckCookie(ctx, session, dto); err != nil {
		return err
	}
	return nil
}

func (p *service) GetPost(ctx context.Context, dto *module.HomePageDTO) error {
	if err := p.post.GetPost(ctx, dto); err != nil {
		return err
	}
	return nil
}
