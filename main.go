package main

import (
	"fmt"
	"todo-app/app/models"
)

func main() {
	//fmt.Println(config.Config)
	//log.Print("test")

	fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@test.com"
		u.Password = "123456"
		fmt.Println(u)

		u.CreateUser()

	*/

	u, _ := models.GetUser(1)
	fmt.Println(u)
}
