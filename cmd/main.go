package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ZangarZaynesh/forum/db"
	"github.com/ZangarZaynesh/forum/internal/composites"
)

func main() {
	ctx := context.Background()
	db := db.CheckDB()

	router := http.NewServeMux()
	User := composites.NewUserComposite(ctx, db)
	Post := composites.NewPostComposite(ctx, db)
	User.Handler.Register(router)
	Post.Handler.Register(router)
	fmt.Println("localhost:8080 listening...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("error lissen and serve main func: ---> %v\n", err)
	}
	defer db.Close()
}
