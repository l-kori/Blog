/*
控制器部分，实现插入文章
*/

package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func insertarticle(c *gin.Context) {
	// Id          int       //文章ID
	// title       string    //文章标题
	// content     string    //文章内容
	// createtime  int64     //创建时间

	title := c.PostForm("titile")
	content := c.PostForm("content")
	username := c.PostForm("username")
	fmt.Println(title, content, username)
	// token_id := c.PostForm("token")

	//检查是否登录，如果登录，则插入
	// err := models.QueryTokenWihtCon(username,token)
	// if err > 0 {
	// 	article := models.Article{0, titile, content, time.Now().Unix()}
	// 	_, err := models.InsertArticle(article)
	// 	if err != nil {
	// 		c.JSON(http.StatusOK, gin.H{"code": 200, "massage": "插入文章成功"})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"code": 0, "massage": "插入文章失败"})
	// 	}
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"code": 0, "massage": "未登录"})
	// }
}
