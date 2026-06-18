package handlers

import "net/http"

func APIInfo(
	w http.ResponseWriter,
	r *http.Request,
) {

	sendJSON(
		w,
		http.StatusOK,
		map[string]string{
			"name":        "APIForge",
			"version":     "1.0",
			"description": "Task Management REST API built with Go",
		},
	)
}
