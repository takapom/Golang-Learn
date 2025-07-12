package service

import (
	"fmt"
	"go_sample/model"
	"go_sample/repository"
)

// サービス層に必要な関数の型定義
type UserService interface {
	GetUser(id int) (*model.User, error)
	RegisterUer(name string, age int) (*model.User, error)
}

// 実際にサービス層の関数を実装させる定義
type userService struct {
	repo repository.UserRepository
}

func (s *userService) GetUser(id uint) (*model.User, error) {
	//そのままレポジトリ層へ渡す(ユーザーを探す操作を)
	return s.repo.FindByID(id)
}

func (s *userService) RegisterUer(name string, age int) (*model.User, error) {
	//バリデーションを行い、成功ならレポジトリに渡す
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	//ユーザー構造に当てはめる
	u := &model.User{Name: name, Age: age}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}
