package composites

import (
	"context"
	"database/sql"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	post3 "github.com/ZangarZaynesh/forum/internal/adapters/handlers/post"
	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/adapters/repository/post"
	"github.com/ZangarZaynesh/forum/internal/domain"
	post2 "github.com/ZangarZaynesh/forum/internal/domain/post"
)

type PostComposite struct {
	Repo    repository.Post
	Service domain.Post
	Handler handlers.Post
}

func NewPostComposite(ctx context.Context, db *sql.DB) *PostComposite {
	repo := post.NewRepository(db)
	service := post2.NewService(repo)
	handler := post3.NewHandler(ctx, service)
	return &PostComposite{
		Repo:    repo,
		Service: service,
		Handler: handler,
	}
}
