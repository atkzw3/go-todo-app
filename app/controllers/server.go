package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo-app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, filename := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", filename))
	}

	// エラーの場合はpanicになるので注意
	templates := template.Must(template.ParseFiles(files...))

	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
