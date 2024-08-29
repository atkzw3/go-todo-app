package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("app/views/templates/top.html")

	data := make(map[string]interface{})
	data["title"] = "テストです！！"

	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
