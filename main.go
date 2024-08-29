package main

import (
	"fmt"
	"todo-app/app/models"
)

func main() {
	//fmt.Println(config.Config)
	//log.Print("test")

	//fmt.Println(models.Db)
	//
	//u := &models.User{}
	//u.Name = "test"
	//u.Email = "test@test.com"
	//u.Password = "123456"
	//fmt.Println(u)
	//
	//u.CreateUser()

	//u, _ := models.GetUser(1)
	//fmt.Println(u)
	//
	//u.Name = "update_name"
	//u.Email = "update_email@test.com"
	//u.Password = models.Encrypt("123123123")
	//
	//fmt.Println(u)
	//
	//u2, _ := models.GetUser(1)
	//fmt.Println(u2)
	//
	//u2.DeleteUser()
	//u3, _ := models.GetUser(1)
	//fmt.Println(u3) // データがないので初期値が返ってくる

	u2, _ := models.GetUser(1)
	fmt.Println(u2)
	c := "コンテンツ！!"

	err := u2.CreateTodo(c)
	if err != nil {
		fmt.Println(err)
	}

	t, _ := models.GetTodo(6)
	fmt.Println("todo", t)

	todoAll, _ := models.GetTodoAll()
	fmt.Println(todoAll)

	for _, v := range todoAll {
		fmt.Println(v)
	}

	fmt.Println("GetByUser メソッド確認")
	todos, err := u2.GetTodosByUser()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range todos {
		fmt.Println(v)
	}
}
