package controllers

import (
	"Blog/server/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	username := c.PostForm("username")
	token_id := c.PostForm("token_id")
	article_id := c.PostForm("article_id")
	//检查token是否存在
	id := models.QueryIdWithToken_id(username, token_id)
	fmt.Println(id)
	if id > 0 {
		token_due_time := models.QueryToken_due_timeWithUsername(username)
		token_due_time_now := time.Now().Unix()
		token_time_count := token_due_time_now - token_due_time
		if token_time_count > 86400 { //86400
			err := models.DeleteTokenWithUsername(username, token_id)
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "用户已退出"})
		} else { //用户登录，且token存在，进行文章删除操作
			err := models.DeleteArticleWithID(article_id)
			fmt.Println(err)
			if err != nil {
				c.JSON(0, gin.H{"massage": "文章删除失败"})
			} else {
				c.JSON(200, gin.H{"massage": "文章删除成功"})
			}
		}
	} else {
		c.JSON(0, gin.H{"massage": "用户已退出"}) //用户token不存在
	}

}
