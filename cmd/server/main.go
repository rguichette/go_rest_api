package main

import (
	"fmt"

	"github.com/rguichette/go_rest_api/internal/db"
)

//Run - is going to be responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")
	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to the Database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	return nil
}

func main() {
	err := Run()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Go Rest API Course")
}
