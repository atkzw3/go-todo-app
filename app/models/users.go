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

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserId    int
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

func (u *User) DeleteUser() (err error) {
	cmdU := `delete from users where id=?`
	_, err = Db.Exec(cmdU, u.ID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email=?`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd := `insert into sessions(
                     uuid, 
                     email, 
                     user_id, 
                     created_at) values(?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id=? and email=?`

	err2 := Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.CreatedAt)

	if err2 != nil {
		log.Println(err2)
	}
	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid=? `

	err = Db.QueryRow(cmd, s.UUID).Scan(
		&s.ID,
		&s.UUID,
		&s.Email,
		&s.UserId,
		&s.CreatedAt)

	if err != nil {
		log.Println(err)
	}

	if s.ID != 0 {
		valid = true
	}
	return valid, err
}
