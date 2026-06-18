package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"apiforge/internal/models"
	"apiforge/internal/storage"
	"apiforge/internal/utils"

	"github.com/go-chi/chi/v5"
)

func GetTasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	statusFilter :=
		r.URL.Query().Get(
			"status",
		)

	search :=
		r.URL.Query().Get(
			"search",
		)

	sortBy :=
		r.URL.Query().Get(
			"sort",
		)
	if statusFilter != "" {

		var filtered []models.Task

		for _, task := range storage.Tasks {

			if task.Status ==
				statusFilter {

				filtered =
					append(
						filtered,
						task,
					)
			}
		}

		if search != "" {

			var filtered []models.Task

			for _, task := range storage.Tasks {

				if strings.Contains(
					strings.ToLower(
						task.Title,
					),
					strings.ToLower(
						search,
					),
				) {

					filtered =
						append(
							filtered,
							task,
						)
				}
			}

			sendJSON(
				w,
				http.StatusOK,
				filtered,
			)

			return
		}

		tasks :=
			make(
				[]models.Task,
				len(storage.Tasks),
			)
		if sortBy == "title" {

			sort.Slice(
				tasks,
				func(i, j int) bool {

					return tasks[i].Title <
						tasks[j].Title
				},
			)
		}
		if sortBy == "priority" {

			order :=
				map[string]int{
					"High":   3,
					"Medium": 2,
					"Low":    1,
				}

			sort.Slice(
				tasks,
				func(i, j int) bool {

					return order[tasks[i].Priority] >
						order[tasks[j].Priority]
				},
			)
		}

		copy(
			tasks,
			storage.Tasks,
		)
		sendJSON(
			w,
			http.StatusOK,
			tasks,
		)

		return
	}

	sendJSON(
		w,
		http.StatusOK,
		storage.Tasks,
	)
}

func GetTaskByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	idParam := chi.URLParam(
		r,
		"id",
	)

	id, err := strconv.Atoi(
		idParam,
	)

	if err != nil {

		sendError(
			w,
			http.StatusBadRequest,
			"Invalid ID",
		)

		return
	}

	for _, task := range storage.Tasks {

		if task.ID == id {

			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			json.NewEncoder(w).Encode(
				task,
			)

			return
		}
	}

	sendError(
		w,
		http.StatusNotFound,
		"Task not found",
	)
}

func CreateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	var task models.Task

	err := json.NewDecoder(
		r.Body,
	).Decode(&task)

	if err != nil {

		sendError(
			w,
			http.StatusBadRequest,
			"Invalid JSON",
		)

		return
	}
	if !utils.IsValidStatus(
		task.Status,
	) {

		sendError(
			w,
			http.StatusBadRequest,
			"Invalid status",
		)

		return
	}

	if !utils.IsValidPriority(
		task.Priority,
	) {

		sendError(
			w,
			http.StatusBadRequest,
			"Invalid priority",
		)

		return
	}

	if len(task.Title) < 3 {
		sendJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "title must contain at least 3 characters",
			},
		)
		return
	}

	task.ID = len(
		storage.Tasks,
	) + 1

	storage.Tasks = append(
		storage.Tasks,
		task,
	)

	storage.SaveTasks()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusCreated,
	)

	json.NewEncoder(w).Encode(
		task,
	)
}

func UpdateTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	idParam := chi.URLParam(
		r,
		"id",
	)

	id, err := strconv.Atoi(
		idParam,
	)

	if err != nil {
		sendError(
			w,
			http.StatusBadRequest,
			"Invalid ID",
		)
		return
	}

	var updatedTask models.Task

	err = json.NewDecoder(
		r.Body,
	).Decode(&updatedTask)

	if err != nil {
		sendError(
			w,
			http.StatusBadRequest,
			"Invalid JSON",
		)
		return
	}
	if !utils.IsValidStatus(
		updatedTask.Status,
	) {

		sendError(
			w,
			http.StatusBadRequest,
			"Invalid status",
		)

		return
	}
	if !utils.IsValidPriority(
		updatedTask.Priority,
	) {

		sendError(
			w,
			http.StatusBadRequest,
			"Invalid priority",
		)

		return
	}
	for i, task := range storage.Tasks {

		if task.ID == id {

			updatedTask.ID = id

			storage.Tasks[i] =
				updatedTask

			storage.SaveTasks()
			sendJSON(
				w,
				http.StatusOK,
				updatedTask,
			)

			return
		}
	}

	sendError(
		w,
		http.StatusNotFound,
		"Task not found",
	)
}

func DeleteTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	idParam := chi.URLParam(
		r,
		"id",
	)

	id, err := strconv.Atoi(
		idParam,
	)

	if err != nil {
		sendError(
			w,
			http.StatusBadRequest,
			"Invalid ID",
		)
		return
	}

	for i, task := range storage.Tasks {

		if task.ID == id {

			storage.Tasks =
				append(
					storage.Tasks[:i],
					storage.Tasks[i+1:]...,
				)

			storage.SaveTasks()

			sendJSON(
				w,
				http.StatusOK,
				map[string]string{
					"message": "Task deleted successfully",
				},
			)

			return
		}
	}

	sendError(
		w,
		http.StatusNotFound,
		"Task not found",
	)
}
