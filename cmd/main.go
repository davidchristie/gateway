//go:generate go run github.com/99designs/gqlgen generate

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/davidchristie/gateway/exec"
	"github.com/davidchristie/gateway/middleware"
	"github.com/davidchristie/gateway/resolvers"
	"github.com/rs/cors"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()

	mux.Handle("/", handler.Playground("GraphQL playground", "/query"))
	mux.Handle("/query", handler.GraphQL(exec.NewExecutableSchema(exec.Config{Resolvers: resolvers.NewRootResolver()})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, cors.Default().Handler(middleware.Middleware(mux))))
}
