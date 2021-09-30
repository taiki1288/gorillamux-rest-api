package repository

import (
	"context"
	"log"

	"../entity"

	"cloud.google.com/go/firestore"
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

const (
	projectId      string = "pragmatic-reviews"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID": post.ID,
		"Title": post.Title,
		"text": post.Text,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}
