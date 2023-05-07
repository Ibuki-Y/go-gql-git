package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Ibuki-Y/go-gql-git/graph"
	"github.com/Ibuki-Y/go-gql-git/graph/services"
	"github.com/Ibuki-Y/go-gql-git/internal"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const (
	defaultPort = "8000"
	dbFile      = "./DB/sqlite.db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbFile))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := services.New(db)

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{
		Resolvers: &graph.Resolver{
			Srv:     service,
			Loaders: graph.NewLoaders(service),
		},
		Complexity: graph.ComplexityConfig(),
	}))
	srv.Use(extension.FixedComplexityLimit(100))

	boil.DebugMode = true

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
