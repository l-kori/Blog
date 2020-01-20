/*
控制器部分，实现业务逻辑
*/

package controllers

import (
	"Blog/server/models"
	"Blog/server/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//}

//处理注册
func RegisterPost(c *gin.Context) {
	//获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	// trades := c.PostForm("trades")
	// fileHeader, err := c.FormFile("head_img")
	// fmt.Println(trades, fileHeader)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUser_idWithUsername(username)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名已经存在"})
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	//如果错误存在
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "注册成功"})
	}

}
