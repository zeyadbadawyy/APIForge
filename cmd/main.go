package main

import (
	"fmt"
	"net/http"

	"apiforge/internal/routes"
)

func main() {
	router := routes.RegisterRoutes()

	fmt.Println(
		"APIForge running on :8080",
	)

	http.ListenAndServe(
		":8080",
		router,
	)
}
