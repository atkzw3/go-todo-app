package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo-app/app/models"
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
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}

// ログイン中か確認
func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie") // authenticateメソッドで指定した
	if err == nil {
		session = models.Session{UUID: cookie.Value}

		if ok, _ := session.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}

	return session, err
}
