package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-app/app/models"
)

func top(w http.ResponseWriter, req *http.Request) {
	_, err := session(w, req)
	// 未ログイン
	if err != nil {
		generateHTML(w, "helllllllo", "layout", "public_nav", "top")
	} else {
		// ログイン済み
		http.Redirect(w, req, "/todos", 302)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	session, err := session(w, req)
	if err != nil {
		http.Redirect(w, req, "/", 302)
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		log.Println(user)

		todos, err2 := user.GetTodosByUser()
		if err2 != nil {
			log.Println(err2)
		}
		user.Todos = todos

		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, req *http.Request) {
	_, err := session(w, req)
	if err != nil {
		http.Redirect(w, req, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, req *http.Request) {
	session, err := session(w, req)
	if err != nil {
		http.Redirect(w, req, "/login", 302)
	} else {
		err := req.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := req.PostFormValue("content")
		if content == "" {
			log.Println("content is empty")
		}

		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, req, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, req *http.Request) {
	_, err := session(w, req)
	if err != nil {
		fmt.Println("ログインしていない")
		http.Redirect(w, req, "/login", 302)
	}

	id := req.PathValue("id")

	fmt.Println("editのid = ", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}

	todo, err := models.GetTodo(idInt)

	if err != nil {
		log.Println(err)
	}

	generateHTML(w, todo, "layout", "private_navbar", "todo_edit")
}
