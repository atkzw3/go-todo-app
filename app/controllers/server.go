package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
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

var validPath = regexp.MustCompile("^/todos/(edit|update)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	fmt.Println("parseURL")

	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)

		fmt.Println("q = ", q)
		if q == nil {
			http.NotFound(w, r)
			return
		}

		// idを取得
		qi, err := strconv.Atoi(q[2])

		fmt.Println("qi = ", qi)
		if err != nil {
			fmt.Println("変換errr")
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)

		fmt.Println("parseURL 完了")
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
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	// / が末尾についていれば、idなどを渡すことが可能
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))

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
