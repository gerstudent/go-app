package main

import (
	"fmt"

	"github.com/gerstudent/go-app/internal/db"
)

func Run() error {
	fmt.Println("start")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate the database")
		return err
	}

	fmt.Println("Successfully connected and pinged database")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}