package handlers

import "net/http"

func HealthCheck(
	w http.ResponseWriter,
	r *http.Request,
) {

	sendJSON(
		w,
		http.StatusOK,
		map[string]string{
			"status": "healthy",
		},
	)
}
