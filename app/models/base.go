package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/daichi-sato-design/todoApp_go_sqlite/config"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
)

func init(){
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil{
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}

func CreateUUID() (uuidobj uuid.UUID){
	uuidobj, _ = uuid.NewV1()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string){
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}