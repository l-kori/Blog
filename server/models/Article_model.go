/*
	Article表，存储文章
*/

package models

import (
	"Blog/server/database"
	"fmt"
)

type Article struct {
	Id        int    //文章ID
	Title     string //文章标题
	Content   string //文章内容
	Ceatetime int64  //创建时间

}

//--------------数据库操作-----------------

//根据文章ID查询内容
func QueryArticleWightID(id string) int {
	sql := fmt.Sprintf("select * from article where id='%s'", id)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id1 := 0
	row.Scan(&id)
	return id1
}

//插入一篇文章
func InsertArticle(article Article) (int64, error) {
	return database.ModifyDB("insert into aerticle(title,content,createtime) value (?,?,?)", article.title, article.content, article.Createtime)
}

//删除文章
func DeleteArticleWithID(id int) (int64, error) {
	return database.ModifyDB("delete from article where id = '%s'", id)
}
