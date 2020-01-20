/*
	model，数据表，以及操作部分
*/

package models

import (
	"Blog/server/database"
	"fmt"
	"time"
)

type User struct {
	Id       int
	Username string
	Password string
	// trades     string
	// head_img   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

//--------------数据库操作-----------------

//插入
func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, time.Now().Unix())
}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUser_idWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUser_idWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}

//根据id查询用户所有数据
func QueryUser_allWithId(id int) string {
	sql := fmt.Sprintf("select username from user where id = '%d", id)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	username := ""
	row.Scan(&username)
	return username
}
