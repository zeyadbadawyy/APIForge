package main

import (
	"fmt"
	"net/http"
	"os"

	"apiforge/internal/routes"
	"apiforge/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	storage.LoadTasks()
	router := routes.RegisterRoutes()

	fmt.Println(
		"APIForge running on :8080",
	)
	godotenv.Load()

	port :=
		os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(
		":"+port,
		router,
	)
}
