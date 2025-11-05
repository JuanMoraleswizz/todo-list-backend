package usecase

import (
	"errors"
	"todo-list/internal/domain"
)

type todoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUseCase(todoRepo domain.TodoRepository) domain.TodoUseCase {
	return &todoUseCase{todoRepo: todoRepo}
}

func (u *todoUseCase) CreateTodo(title, description string) (*domain.Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	todo := &domain.Todo{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	err := u.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *todoUseCase) GetAllTodos() ([]domain.Todo, error) {
	return u.todoRepo.GetAll()
}

func (u *todoUseCase) GetTodoByID(id uint) (*domain.Todo, error) {
	return u.todoRepo.GetByID(id)
}

func (u *todoUseCase) UpdateTodo(id uint, title, description string, completed *bool) (*domain.Todo, error) {
	todo, err := u.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		todo.Title = title
	}
	if description != "" {
		todo.Description = description
	}
	if completed != nil {
		todo.Completed = *completed
	}

	err = u.todoRepo.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *todoUseCase) DeleteTodo(id uint) error {
	return u.todoRepo.Delete(id)
}