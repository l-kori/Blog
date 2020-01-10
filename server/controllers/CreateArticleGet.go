package controllers

import (
	"Blog/server/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//处理新增文章
func CreateArticle(c *gin.Context) {
	username := c.PostForm("username")
	token_id := c.PostForm("token_id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	//检查token是否存在
	token_due_time := models.QueryToken_due_timeWithUsername(username)
	token_due_time_now := time.Now().Unix()
	token_time_count := token_due_time_now - token_due_time
	if token_time_count > 86400 { //86400
		err := models.DeleteTokenWithUsername(username, token_id)
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "用户已退出"})
	} else { //用户登录，且token存在，进行文章插入操作
		article := models.Article{title, content, time.Now().Unix()}
		_, err := models.InsertArticle(article)
		if err != nil { //存在错误
			c.JSON(0, gin.H{"massage": "失败"}) //插入文章失败
		} else {
			c.JSON(200, gin.H{"massage": "成功"}) //插入文章成功
		}
	}

}
