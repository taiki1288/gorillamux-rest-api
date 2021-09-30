package repository

type PostRepository interface {
	Save(post *Post) (*Post, error)
	FindAll() ([]Post, error)
}