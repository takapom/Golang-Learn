package service

import (
	"Todo_layered/model"
	"Todo_layered/repository"
)

// typeの定義
//①service層で何をできるかを明確化
//②リポジトリ層への橋渡し

type TodoService interface {
	CreateTodo(title, description string) (*model.Todo, error)
	GetTodos() ([]model.Todo, error)
	GetTodo(id uint) (*model.Todo, error)
	UpdateTodo(id uint, title, description string, completed bool) (*model.Todo, error)
	DeleteTodo(id uint) error
}

// リポジトリ層に渡すため
type todoService struct {
	repo repository.TodoRepository
}

// todoServieのrepoを明確化する
func NewTodoService(repo repository.TodoRepository) TodoService {
}

func (s *todoService) CreateTodo(title, description string) (*model.Todo, error) {
	todo := &model.Todo{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodo

func (s *todoService) GetTodo(id uint) (*model.Todo, error) {
	return s.GetByID(id)
}
