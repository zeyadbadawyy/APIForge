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
	inProgress := 0

	for _, task := range storage.Tasks {

		switch task.Status {

		case "Completed":
			completed++

		case "Pending":
			pending++

		case "In Progress":
			inProgress++
		}
	}

	sendJSON(
		w,
		http.StatusOK,
		map[string]int{
			"totalTasks": total,
			"completed":  completed,
			"pending":    pending,
			"inProgress": inProgress,
		},
	)
}
