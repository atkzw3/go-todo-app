package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (content, user_id, created_at) values(?, ?, ?)`

	log.Println(cmd, content, u)

	_, err = Db.Exec(cmd, content, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select * from todos where id = ?`

	log.Println(cmd, id)

	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	if err != nil {
		log.Fatalln(err)
	}

	return todo, err
}

func GetTodoAll() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}
	cErr := rows.Close()
	if cErr != nil {
		return nil, cErr
	}

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	cErr := rows.Close()
	if cErr != nil {
		return nil, cErr
	}
	return todos, err
}
