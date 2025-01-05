package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/OsqY/GoingNext/internal/config"
	"github.com/OsqY/GoingNext/internal/db"
	"github.com/OsqY/GoingNext/internal/http_internal"
	"github.com/OsqY/GoingNext/internal/http_internal/handlers"
	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, fmt.Sprintf("host=%s port=%v user=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.DBName))
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}

	queries := db.New(conn)

	userHandler := handlers.NewUserHandler(queries)
	authHandler := handlers.NewAuthHandler(queries)
	roleHandler := handlers.NewRoleHandler(queries)
	fileHandler := handlers.NewFileHandler(config)

	router := http_internal.NewRouter(userHandler, authHandler, roleHandler, fileHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}
