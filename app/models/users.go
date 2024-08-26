package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmdU := `insert into users(
                  uuid, 
                  name, 
                  email, 
                  password, 
                  created_at) values(?, ?, ?, ?, ?)`

	_, err = Db.Exec(
		cmdU,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	if err != nil {
		log.Println(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id=?`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmdU := `update users set name=?, email=?, password=? where id=?`

	_, err = Db.Exec(cmdU, u.Name, u.Email, u.Password, u.ID)

	if err != nil {
		log.Println(err)
	}
	return err
}
