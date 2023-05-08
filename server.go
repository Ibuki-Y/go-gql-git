package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Ibuki-Y/go-gql-git/graph"
	"github.com/Ibuki-Y/go-gql-git/graph/services"
	"github.com/Ibuki-Y/go-gql-git/internal"
	"github.com/Ibuki-Y/go-gql-git/middlewares/auth"
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
		Directives: graph.Directive,
	}))
	srv.Use(extension.FixedComplexityLimit(100))

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		log.Println("before OperationHandler")
		res := next(ctx)
		defer log.Println("after OperationHandler")
		return res
	})
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		log.Println("before ResponseHandler")
		res := next(ctx)
		defer log.Println("after ResponseHandler")
		return res
	})
	srv.AroundRootFields(func(ctx context.Context, next graphql.RootResolver) graphql.Marshaler {
		log.Println("before RootResolver")
		res := next(ctx)
		defer func() {
			var b bytes.Buffer
			res.MarshalGQL(&b)
			log.Println("after RootResolver", b.String())
		}()
		return res
	})
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		res, err = next(ctx)
		log.Println(res)
		return
	})

	boil.DebugMode = true

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.AuthMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
