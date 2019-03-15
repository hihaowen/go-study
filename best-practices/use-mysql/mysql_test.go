package use_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	sq "github.com/Masterminds/squirrel"
)

var Db *sql.DB
var dsn = "root:root@tcp(127.0.0.1:3306)/test"

func Connect() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)

	Db = db
}

func TestSelect(t *testing.T) {
	// 连接数据库
	Connect()

	users := sq.Select("id, user_name").From("user")

	// users = users.Where(sq.Eq{"user_name": "George"})
	users = users.Limit(1)

	rows, err := users.RunWith(Db).Query()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id   int64
			user_name string
		)
		if err := rows.Scan(&id, &user_name); err != nil {
			continue
		}
		fmt.Printf("id %d name is %s\n", id, user_name)
	}
}
