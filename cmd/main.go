package main

import (
	"fmt"
	"net/http"
)

func homeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(
		w,
		"Welcome to APIForge",
	)
}

func main() {
	http.HandleFunc(
		"/",
		homeHandler,
	)

	fmt.Println(
		"Server running on :8080",
	)

	http.ListenAndServe(
		":8080",
		nil,
	)
}
