package storage

import "apiforge/internal/models"

var Tasks = []models.Task{
	{
		ID:          1,
		Title:       "Learn Go",
		Description: "Study structs and slices",
		Status:      "In Progress",
		Priority:    "High",
	},
	{
		ID:          2,
		Title:       "Build APIForge",
		Description: "Create REST API",
		Status:      "Pending",
		Priority:    "Medium",
	},
}