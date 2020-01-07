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

//查询token记录
func QueryTokenWihtCon(username string, token_id string) int {
	sql := fmt.Sprintf("select username from token where username = '%s' and token_id = '%s'", username, token_id)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//查询token的有效期
func QueryToken_due_timeWithusername(username string) int64 {
	sql := fmt.Sprintf("select token_due_time from token where username = '%s'", username)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	var token_due_time int64
	row.Scan(&token_due_time)
	return token_due_time
}

//删除一条记录
func DeleteTokenWithUsername(username string, token string) error {
	sql := fmt.Sprintf("delete from token where username = '%s' and token_id = '%s'", username, token)
	fmt.Println(sql)
	_, err := database.ModifyDB(sql)
	return err

}
