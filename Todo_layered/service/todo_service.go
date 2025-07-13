// サービス層は「どうやってデータを保存／取得しているか」を知らずに、ただ「リポジトリに任せる」だけにしたい。
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
	return &todoService{repo: repo}
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

func (s *todoService) GetTodos() ([]model.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) GetTodo(id uint) (*model.Todo, error) {
	return s.repo.GetByID(id)
}

// Todoの更新処理
func (s *todoService) UpdateTodo(id uint, title, description string, completed bool) (*model.Todo, error) {
	// まずは該当するidのタスク取得
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	//取得したidを更新する処理
	if title != "" {
		todo.Title = title
	}
	if description != "" {
		todo.Description = description
	}
	todo.Completed = completed

	//上の情報をリポジトリ層に渡す
	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}
