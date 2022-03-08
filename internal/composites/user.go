package composites

import (
	"context"
	"database/sql"

	"github.com/ZangarZaynesh/forum/internal/adapters/handlers"
	user3 "github.com/ZangarZaynesh/forum/internal/adapters/handlers/user"
	"github.com/ZangarZaynesh/forum/internal/adapters/repository"
	"github.com/ZangarZaynesh/forum/internal/adapters/repository/user"
	"github.com/ZangarZaynesh/forum/internal/domain"
	user2 "github.com/ZangarZaynesh/forum/internal/domain/user"
)

type UserComposite struct {
	Repo    repository.User
	Service domain.User
	Handler handlers.User
}

func NewUserComposite(ctx context.Context, db *sql.DB) *UserComposite {
	repo := user.NewRepository(db)
	service := user2.NewService(repo)
	handler := user3.NewHandler(ctx, service)
	return &UserComposite{
		Repo:    repo,
		Service: service,
		Handler: handler,
	}
}
