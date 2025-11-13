package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Egor332/TokenTransferApi/database"
	"github.com/Egor332/TokenTransferApi/graph"
	"github.com/Egor332/TokenTransferApi/repository"
	"github.com/Egor332/TokenTransferApi/service"
)

func main() {
	database.Connect()

	log.Println("Database connection is ready")

	repo := repository.NewGormWalletRepository()
	walletService := service.NewWalletTransferService(repo, database.DB)

	resolver := &graph.Resolver{
		WalletService: walletService,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Server running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Println("Server started")
}
