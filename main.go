package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/todo"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}
	repo, err := todo.NewRepo(connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	handler := todo.NewHandler(repo)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
