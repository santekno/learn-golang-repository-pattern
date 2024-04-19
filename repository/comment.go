package repository

import (
	"context"
	"fmt"

	model "github.com/santekno/learn-golang-repository-pattern/model"
)

func (repo *CommentRepo) Insert(ctx context.Context, comment model.Comment) (model.Comment, error) {
	result, err := repo.DB.ExecContext(ctx, "INSERT INTO comments(email,comment) VALUES(?,?)", comment.Email, comment.Email)
	if err != nil {
		return comment, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(insertId)
	return comment, nil
}

func (repo *CommentRepo) FindById(ctx context.Context, id int32) (model.Comment, error) {
	var comment model.Comment
	query := "SELECT id, email, comment FROM comments WHERE id=? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}
	}
	return comment, nil
}

func (repo *CommentRepo) FindAll(ctx context.Context) ([]model.Comment, error) {
	var comments []model.Comment
	query := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			fmt.Printf("error scan rows %v", err)
			continue
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
