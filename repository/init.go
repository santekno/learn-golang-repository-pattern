package repository

import (
	"context"
	"database/sql"

	model "github.com/santekno/learn-golang-repository-pattern/model"
)

type CommentRepo struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepo{
		DB: db,
	}
}

type CommentRepository interface {
	Insert(ctx context.Context, comment model.Comment) (model.Comment, error)
	FindById(ctx context.Context, id int32) (model.Comment, error)
	FindAll(ctx context.Context) ([]model.Comment, error)
}
