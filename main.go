package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/santekno/learn-golang-repository-pattern/model"
	repository "github.com/santekno/learn-golang-repository-pattern/repository"
)

func main() {
	ctx := context.Background()
	db := GetConnection()
	defer db.Close()

	commentRepo := repository.NewCommentRepository(db)

	// find all data comments
	comments, err := commentRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, cm := range comments {
		fmt.Printf("data %d: %v\n", cm.Id, cm)
	}

	// find all data by id
	comment, err := commentRepo.FindById(ctx, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data : %v", comment)

	// insert data
	id, err := commentRepo.Insert(ctx, model.Comment{Email: "test@gmail.com", Comment: "komentar yuk"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("lastId: %v", id)
}

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:belajargolang@tcp(localhost:3306)/belajar-golang")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
