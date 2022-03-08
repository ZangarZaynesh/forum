package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ZangarZaynesh/forum/db"
	user3 "github.com/ZangarZaynesh/forum/internal/adapters/handlers/user"
	"github.com/ZangarZaynesh/forum/internal/adapters/repository/user"
	user2 "github.com/ZangarZaynesh/forum/internal/domain/user"
)

func main() {
	ctx := context.Background()
	db := db.CheckDB()

	router := http.NewServeMux()

	userRepository := user.NewRepository(db)
	userService := user2.NewService(userRepository)
	userHandler := user3.NewHandler(ctx, userService)
	userHandler.Register(router)
	fmt.Println("localhost:8080 listening...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("error lissen and serve main func: ---> %v\n", err)
	}
	defer db.Close()
}
