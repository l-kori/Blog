/*
操作数据库，mysql
*/

package database

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func Initmysql() {
	if db == nil {
		db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Blog?charset=utf8")
		fmt.Println(db)
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
