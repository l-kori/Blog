package models

import (
	"Blog/server/database"
	"fmt"
)

type Token struct {
	Username       string
	Token_id       string
	Token_due_time int64
}

//按条件查询id
func QueryTokenWightCon(con string) int {
	sql := fmt.Sprintf("select id from token %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名和密码，查询id
func QueryTokenWithParam(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryTokenWightCon(sql)
}

//插入一条token记录
func InsertToken(token Token) (int64, error) {
	return database.ModifyDB("insert into token(username,token_id,token_due_time) values (?,?,?)",
		token.Username, token.Token_id, token.Token_due_time)
}

//根据用户名查询id
func QueryTokenWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//查询token的有效期
func QueryToken_due_timeWithUsername(username string) int64 {
	sql := fmt.Sprintf("select token_due_time from token where username = '%s'", username)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	var token_due_time int64
	row.Scan(&token_due_time)
	return token_due_time
}

//按条件查询token值
func QueryToken_idWightCon(con string) string {
	sql := fmt.Sprintf("select token_id from token %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	token_id := ""
	row.Scan(&token_id)
	return token_id
}

//查询token_id
func QueryToken(username string) string {
	sql := fmt.Sprintf("where username = '%s'", username)
	return QueryToken_idWightCon(sql)
}

//修改token有效期
func AlterToken_due_timeWhthUsername(token_due_time int64, username string) error {
	sql := fmt.Sprintf("update token set token_due_time = '%d' where username = '%s'", token_due_time, username)
	fmt.Println(sql)
	_, err := database.ModifyDB(sql)
	return err
}

//删除一条记录
func DeleteTokenWithUsername(username string, token string) error {
	sql := fmt.Sprintf("delete from token where username = '%s' and token_id = '%s'", username, token)
	fmt.Println(sql)
	_, err := database.ModifyDB(sql)
	return err

}
