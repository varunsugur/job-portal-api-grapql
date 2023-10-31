package main

import (
	"context"
	"fmt"
	database "golang/databse"
	"golang/graph"
	"golang/repository"
	"golang/service"

	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog/log"
)

const defaultPort = "8080"

func main() {
	svc, err := StartApp()
	if err != nil {
		log.Info().Err(err).Msg("error in Start App")

	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Svc: svc}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil))
}

func StartApp() (service.UserService, error) {

	db, err := database.Open()
	if err != nil {
		return &service.Service{}, fmt.Errorf("error in connecting database:%w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return &service.Service{}, fmt.Errorf("error in getting databse Instance:%w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return &service.Service{}, fmt.Errorf("error in connecting to database: %w", err)
	}
	repo, err := repository.NewRepo(db)
	if err != nil {
		return &service.Service{}, fmt.Errorf("repository not initiated: %w", err)
	}

	svc, err := service.NewService(repo)
	if err != nil {
		return &service.Service{}, fmt.Errorf("service is not initiated: %w", err)
	}
	return svc, nil

}
