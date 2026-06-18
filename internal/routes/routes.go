package routes

import (
	"net/http"

	"apiforge/internal/handlers"
	"apiforge/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() http.Handler {

	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
	)

	r.Route(
		"/api/v1",
		func(r chi.Router) {

			r.Get(
				"/health",
				handlers.HealthCheck,
			)

			r.Get(
				"/tasks",
				handlers.GetTasks,
			)

			r.Get(
				"/tasks/{id}",
				handlers.GetTaskByID,
			)

			r.Post(
				"/tasks",
				handlers.CreateTask,
			)

			r.Put(
				"/tasks/{id}",
				handlers.UpdateTask,
			)

			r.Delete(
				"/tasks/{id}",
				handlers.DeleteTask,
			)

			r.Get(
				"/stats",
				handlers.GetStats,
			)
		},
	)

	return r
}
