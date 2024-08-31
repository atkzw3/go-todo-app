package controllers

import (
	"log"
	"net/http"
	"todo-app/app/models"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_nav", "signup")
	} else if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_nav", "login")
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := models.GetUserByEmail(r.PostFormValue("email"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		if user.Password == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession()
			if err != nil {
				log.Println(err)
			}

			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

	}
}
