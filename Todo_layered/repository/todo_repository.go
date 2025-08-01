package repository

import (
	"Todo_layered/model"
	"fmt"

	"gorm.io/gorm"
)

//リポジトリ層を使うために
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

// リポジトリ層で使用するメソッド(理想)
type TodoRepository interface {
	Create(todo *model.Todo) error
	GetAll() ([]model.Todo, error)
	GetByID(id uint) (*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id uint) error
}

// 理想を実現するための構造
type todoRepository struct {
	db *gorm.DB
}

// dbに新規で作成するrepo
func (r *todoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

// 全てのTodoを取得する
func (r *todoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	if err := r.db.Find(&todos); err != nil {
		fmt.Println("エラー発生")
	}
	return todos, nil
}

func (r *todoRepository) GetByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&model.Todo{}, id).Error
}
