package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"todo-app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatal(err)
	}

	comU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	uuid STRING UNIQUE NOT NULL,
    	name STRING,
    	email STRING,
    	password STRING,
    	created_at DATETIME,)`, tableNameUser)

	_, err = Db.Exec(comU)
}

func createUUID() (id uuid.UUID) {
	id = uuid.New()
	return id
}

func Encrypt(plane string) (crypt string) {
	crypt = fmt.Sprintf("%x", sha1.Sum([]byte(plane)))
	return crypt
}
