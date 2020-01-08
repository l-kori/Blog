package controllers

import (
	"Blog/server/models"
	"Blog/server/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token_id := c.PostForm("token")
	fmt.Println(username, password, token_id)
	article := models.QueryArticle_content()
	//如果没有传token，则判断为游客浏览模式，只展示文章
	if token_id == "" || username == "" {
		c.JSON(200, gin.H{"massage": "成功", "data": article})
	} else { //用户登录情况下,验证密码
		err := models.QueryUser_idWithParam(username, utils.MD5(password))
		if err > 0 {
			token_due_time := models.QueryToken_due_timeWithUsername(username)
			token_due_time_now := time.Now().Unix()
			token_time_count := token_due_time_now - token_due_time
			if token_time_count > 86400 { //86400
				err := models.DeleteTokenWithUsername(username, token_id)
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{"code": 200, "message": "用户已退出"})
			} else {
				c.JSON(200, gin.H{"massage": "成功", "data": article, "username": username, "token": token_id})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
		}
	}
}
