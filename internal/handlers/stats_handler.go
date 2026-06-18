package handlers

import (
	"net/http"

	"apiforge/internal/storage"
)

func GetStats(
	w http.ResponseWriter,
	r *http.Request,
) {

	total := len(
		storage.Tasks,
	)

	completed := 0
	pending := 0

	for _, task := range storage.Tasks {

		if task.Status ==
			"Completed" {

			completed++
		}

		if task.Status ==
			"Pending" {

			pending++
		}
	}

	sendJSON(
		w,
		http.StatusOK,
		map[string]int{
			"totalTasks": total,
			"completed":  completed,
			"pending":    pending,
		},
	)
}
