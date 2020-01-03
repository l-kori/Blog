package controllers

import (
	"Blog/server/models"
	"Blog/server/utils"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func LoginPost(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	id := UniqueId()
	token_id := username+id
	//判断账号密码是否正确 如果不存在，返回0，如果存在，返回其他
	err  :=models.QueryUserWithParam(username,utils.MD5(password))
	err1 :=models.QueryTokenWightCon(username)
	fmt.Println(err1)
	//判断账号密码是否正确
	if err > 0{
		//判断是否存在token
		if err1 > 0{
			fmt.Println("-------Token存在")
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "Token存在"})
		}else {
			//插入
			fmt.Println("------Token不存在")
			token := models.Token{username,token_id}
			_ ,err := models.InsertToken(token)
			if err != nil{
				c.JSON(http.StatusOK, gin.H{"code": 0, "message": "插入失败"})
			}else {
				c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功","Token":token_id})
			}
		}
	}else {
		c.JSON(http.StatusOK,gin.H{"code":200,"message":"登录失败"})
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
