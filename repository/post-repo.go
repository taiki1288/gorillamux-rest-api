package repository

import (
	"../entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	// entityフォルダの構造体をアクセスしている。
}

type repo struct {}

func NewPostRepository() PostRepository {
	return &repo{}
}