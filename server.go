package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JazzXP/proof-viewer-server/graph"
	"github.com/JazzXP/proof-viewer-server/graph/generated"

	"github.com/jackc/pgx/v4/pgxpool"
)

const defaultPort = "8090"

var pool *pgxpool.Pool

func main() {
	pgHost := os.Getenv("DB_HOST")
	if pgHost == "" {
		pgHost = "localhost"
	}
	pgPort := os.Getenv("DB_PORT")
	if pgPort == "" {
		pgPort = "5432"
	}

	pgUser := os.Getenv("POSTGRES_USER")
	if pgUser == "" {
		panic(fmt.Errorf("POSTGRES_USER must be specified"))
	}

	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	if pgPassword == "" {
		panic(fmt.Errorf("POSTGRES_PASSWORD must be specified"))
	}

	pgDb := os.Getenv("POSTGRES_DB")
	if pgDb == "" {
		panic(fmt.Errorf("POSTGRES_DB must be specified"))
	}

	pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgHost, pgPort, pgUser, pgPassword, pgDb))
	if err != nil {
		log.Print(err)
		panic("Unable to create connection pool")
	}

	fmt.Println("Successfully connected to DB!")
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Print(err)
		panic("Unable to acquire connection")
	}
	err = conn.Conn().Ping(context.Background())
	if err != nil {
		log.Print(err)
		panic("Error with connection")
	}
	conn.Release()
	// GraphQL
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/health", statusHandler(pool))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// HealthStatus return structure for healthcheck
type HealthStatus struct {
	Postgres string `json:"postgres"`
}

func statusHandler(pgPool *pgxpool.Pool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		health := HealthStatus{
			Postgres: "Ok",
		}
		conn, err := pgPool.Acquire(context.Background())
		if err != nil {
			health.Postgres = "Unable to connect"
		}
		defer conn.Release()

		err = conn.Conn().Ping(context.Background())
		if err != nil {
			health.Postgres = "Unable to connect"
		}

		bytes, err := json.MarshalIndent(health, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}
