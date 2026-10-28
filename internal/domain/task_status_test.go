package domain

import (
	"encoding/json"
	"testing"
)

type CreateTaskRequest struct {
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}

func TestCreateTaskRequestWithValidStatus(t *testing.T) {
	body := []byte(`{"title": "Test Task", "status": "todo"}`)

	var req CreateTaskRequest

	err := json.Unmarshal(body, &req)
	if err != nil {
		t.Fatalf("No esperaba este error: %v", err)
	}

	if req.Status != TaskStatusTodo {
		t.Fatalf("Esperaba status todo, recibí: %v", req.Status)
	}
}

func TestCreateTaskRequestWithInvalidStatus(t *testing.T) {
	body := []byte(`{"title": "Test Task", "status": "invalid"}`)

	var req CreateTaskRequest

	err := json.Unmarshal(body, &req)
	if err == nil {
		t.Fatalf("Esperaba error pero recibí nil")
	}

	expectedError := `estado de tarea invalido: "invalid"`

	if err.Error() != expectedError {
		t.Fatalf("Esperaba error: %q, recibí: %q", expectedError, err.Error())
	}
}
