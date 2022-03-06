package main

import (
	"context"
	"database/sql"
	"net/http"

	user3 "github.com/ZangarZaynesh/forum/internal/adapters/handlers/user"
	"github.com/ZangarZaynesh/forum/internal/adapters/repository/user"
	user2 "github.com/ZangarZaynesh/forum/internal/domain/user"
)

func main() {
	ctx := context.Background()
	var db *sql.DB

	router := http.NewServeMux()

	userRepository := user.NewRepository(db)
	userService := user2.NewService(userRepository)
	userHandler := user3.NewHandler(ctx, userService)
	userHandler.Register(router)
}
