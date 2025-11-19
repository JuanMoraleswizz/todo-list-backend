package domain

import (
	"testing"
	"time"
)

func TestTodoStruct(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if todo.ID != 1 {
		t.Errorf("Expected ID 1, got %d", todo.ID)
	}
	if todo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got %s", todo.Title)
	}
	if todo.Completed {
		t.Error("Expected completed to be false")
	}
}