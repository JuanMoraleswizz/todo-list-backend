package usecase

import (
	"errors"
	"testing"
	"todo-list/internal/domain"
)

type mockTodoRepository struct {
	todos []domain.Todo
	err   error
}

func (m *mockTodoRepository) Create(todo *domain.Todo) error {
	if m.err != nil {
		return m.err
	}
	todo.ID = uint(len(m.todos) + 1)
	m.todos = append(m.todos, *todo)
	return nil
}

func (m *mockTodoRepository) GetAll() ([]domain.Todo, error) {
	return m.todos, m.err
}

func (m *mockTodoRepository) GetByID(id uint) (*domain.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, todo := range m.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockTodoRepository) Update(todo *domain.Todo) error {
	if m.err != nil {
		return m.err
	}
	for i, t := range m.todos {
		if t.ID == todo.ID {
			m.todos[i] = *todo
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockTodoRepository) Delete(id uint) error {
	if m.err != nil {
		return m.err
	}
	for i, todo := range m.todos {
		if todo.ID == id {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func TestCreateTodo(t *testing.T) {
	repo := &mockTodoRepository{}
	uc := NewTodoUseCase(repo)

	todo, err := uc.CreateTodo("Test Todo", "Test Description")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got %s", todo.Title)
	}
}

func TestCreateTodoEmptyTitle(t *testing.T) {
	repo := &mockTodoRepository{}
	uc := NewTodoUseCase(repo)

	_, err := uc.CreateTodo("", "Test Description")
	if err == nil {
		t.Error("Expected error for empty title")
	}
}

func TestGetAllTodos(t *testing.T) {
	repo := &mockTodoRepository{
		todos: []domain.Todo{{ID: 1, Title: "Test"}},
	}
	uc := NewTodoUseCase(repo)

	todos, err := uc.GetAllTodos()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}
}

func TestUpdateTodo(t *testing.T) {
	repo := &mockTodoRepository{
		todos: []domain.Todo{{ID: 1, Title: "Original", Completed: false}},
	}
	uc := NewTodoUseCase(repo)

	completed := true
	todo, err := uc.UpdateTodo(1, "Updated", "", &completed)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todo.Title != "Updated" || !todo.Completed {
		t.Error("Todo not updated correctly")
	}
}