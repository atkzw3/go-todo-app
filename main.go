package main

import (
	"fmt"
	"log"
	"todo-app/app/models"
	"todo-app/config"
)

func main() {
	fmt.Println(config.Config)
	log.Print("test")

	fmt.Println(models.Db)
}
