package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/internal/domain"

	"github.com/gorilla/mux"
)

type mockTodoUseCase struct {
	todos []domain.Todo
	err   error
}

func (m *mockTodoUseCase) CreateTodo(title, description string) (*domain.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	todo := &domain.Todo{
		ID:          uint(len(m.todos) + 1),
		Title:       title,
		Description: description,
		Completed:   false,
	}
	m.todos = append(m.todos, *todo)
	return todo, nil
}

func (m *mockTodoUseCase) GetAllTodos() ([]domain.Todo, error) {
	return m.todos, m.err
}

func (m *mockTodoUseCase) GetTodoByID(id uint) (*domain.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, todo := range m.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, m.err
}

func (m *mockTodoUseCase) UpdateTodo(id uint, title, description string, completed *bool) (*domain.Todo, error) {
	return nil, m.err
}

func (m *mockTodoUseCase) DeleteTodo(id uint) error {
	return m.err
}

func TestGetTodos(t *testing.T) {
	mockUC := &mockTodoUseCase{
		todos: []domain.Todo{{ID: 1, Title: "Test Todo"}},
	}
	handler := NewTodoHandler(mockUC)

	req := httptest.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()

	handler.GetTodos(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestCreateTodo(t *testing.T) {
	mockUC := &mockTodoUseCase{}
	handler := NewTodoHandler(mockUC)

	body := map[string]string{
		"title":       "New Todo",
		"description": "Description",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateTodo(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestGetTodo(t *testing.T) {
	mockUC := &mockTodoUseCase{
		todos: []domain.Todo{{ID: 1, Title: "Test Todo"}},
	}
	handler := NewTodoHandler(mockUC)

	req := httptest.NewRequest("GET", "/todos/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handler.GetTodo(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}