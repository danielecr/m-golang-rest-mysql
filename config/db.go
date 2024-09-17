package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Db() *sql.DB {
	// get user from environment
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = "root"
	}
	// get passwd from environment
	passwd := os.Getenv("MYSQL_PASSWORD")
	if passwd == "" {
		passwd = "root"
	}
	db, err := sql.Open("mysql", user+":"+passwd+"@/golang")
	if err != nil {
		panic(err)
	}
	return db
}

func init() {
	DB = Db()
}
