package models

import (
	"Blog/server/database"
	"fmt"
)

type Token struct {
	Username string
	Token_id string
}


//插入一条token记录
func InsertToken(token Token) (int64, error) {
	return database.ModifyDB("insert into token(username,token_id) values (?,?)",
		token.Username,token.Token_id)
}



//根据用户名查询id
func QueryTokenWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//查询token记录
func QueryTokenWightCon(con string) int {
	sql := fmt.Sprintf("select username from token where username = '%s'", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}
