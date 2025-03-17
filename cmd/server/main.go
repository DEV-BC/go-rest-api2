package main

import (
	"context"
	"fmt"
	"github.com/DEV-BC/go-rest-api2/internal/comment"
	"github.com/DEV-BC/go-rest-api2/internal/database"
)

//Run is going to be responsible for the instantiation and startup of the application

func Run() error {
	fmt.Println("Starting up our application ")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.CreateComment(
		context.Background(),
		comment.Comment{
			ID:     "71c5d074-b6cf-11ec-b909-0242ac120002",
			Slug:   "manuel-test",
			Body:   "testing create method",
			Author: "Hello from DevBC",
		},
	)
	fmt.Println(cmtService.GetComment(context.Background(), "71c5d074-b6cf-11ec-b909-0242ac120002"))

	return nil
}

func main() {
	fmt.Println("GO REST API!!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
