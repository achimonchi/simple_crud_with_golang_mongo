package main

import (
	"fmt"

	"github.com/achimonchi/belajar_crud_mongodb/config"
)

func main() {
	fmt.Println("CRUD with Go and Mongo")

	_, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}
}
