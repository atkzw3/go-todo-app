package controllers

import (
	"log"
	"net/http"
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
