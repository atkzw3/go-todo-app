package controllers

import (
	"errors"
	"log"
	"net/http"
	"todo-app/app/models"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_nav", "signup")
		} else {
			// ログイン済み
			http.Redirect(w, r, "/todos", 302)
		}
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
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_nav", "login")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
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

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if !errors.Is(err, http.ErrNoCookie) {
		session := models.Session{UUID: cookie.Value}
		err := session.DeleteSessionByUUID()
		if err != nil {
			log.Println(err)
		}
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
