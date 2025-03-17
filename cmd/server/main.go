package main

import (
	"fmt"
	"github.com/DEV-BC/go-rest-api2/internal/comment"
	transportHTTP "github.com/DEV-BC/go-rest-api2/internal/controller/http"
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

	httpHandler := transportHTTP.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("GO REST API!!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
