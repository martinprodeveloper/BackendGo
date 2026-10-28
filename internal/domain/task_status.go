// Package domain define the domain models and logic for the application.
package domain

import (
	"encoding/json"
	"fmt"
)

// TaskStatus represents the status of a task.
type TaskStatus string

// Valid task statuses.
const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

// IsValid checks if the task status is valid.
func (status TaskStatus) IsValid() bool {
	switch status {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		return true
	default:
		return false
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface for TaskStatus.
func (status *TaskStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	statusTask := TaskStatus(raw)
	if !statusTask.IsValid() {
		return fmt.Errorf("estado de tarea invalido: %q", raw)
	}

	*status = statusTask
	return nil
}
