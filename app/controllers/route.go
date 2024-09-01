package controllers

import (
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
	_, err := session(w, req)
	if err != nil {
		http.Redirect(w, req, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
