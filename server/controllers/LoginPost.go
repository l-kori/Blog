package controllers

import (
	"Blog/server/models"
	"Blog/server/utils"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//判断账号密码是否正确 如果不存在，返回0，如果存在，返回其他
	err := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("222222222222222222222222", err)
	if err > 0 {
		fmt.Println("账号存在")
		err1 := models.QueryTokenWithParam(username)
		fmt.Println("33333333333333333333333333333333", err1)
		if err1 > 0 {
			fmt.Println("token存在")
			models.AlterToken_due_timeWhthUsername(time.Now().Unix(), username) //登录修改token有效期
			models.QueryToken(username)
			token_id := models.QueryToken(username)
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功", "Token": token_id})
		} else {
			fmt.Println("token不存在")
			id := UniqueId()
			token_id := username + id
			token := models.Token{username, token_id, time.Now().Unix()}
			_, err := models.InsertToken(token)
			if err != nil {
				fmt.Println("token插入失败")
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功", "Token": token_id})
			}
		}
	} else {
		fmt.Println("账号不存在")
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录失败"})
	}
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
