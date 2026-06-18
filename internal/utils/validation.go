package utils

var ValidStatuses = map[string]bool{
	"Pending":     true,
	"In Progress": true,
	"Completed":   true,
}

func IsValidStatus(
	status string,
) bool {
	return ValidStatuses[status]
}

var ValidPriorities = map[string]bool{
	"Low":    true,
	"Medium": true,
	"High":   true,
}

func IsValidPriority(
	priority string,
) bool {

	return ValidPriorities[priority]
}
