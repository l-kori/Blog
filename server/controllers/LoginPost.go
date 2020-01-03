package controllers

import (
	"Blog/server/models"
	"Blog/server/utils"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const(
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin = "请重新登陆"
)

func LoginPost(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	claims := &JWTClaims{
		UserID:      1,
		Username:    username,
		Password:    password,
		FullName:    username,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken,err:=getToken(claims)
	if err!=nil{
		c.String(http.StatusNotFound, err.Error())
		return
	}

	//判断账号密码是否正确 如果不存在，返回0，如果存在，返回其他
	err1 :=models.QueryUserWithParam(username,utils.MD5(password))
	if err1 > 0{
		c.JSON(http.StatusOK,gin.H{"code":200,"message":"登录成功","token":signedToken})

	}else {
		c.JSON(http.StatusOK,gin.H{"code":200,"message":"登录失败"})
	}

}
type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      int      `json:"user_id"`
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)


func verifyAction(strToken string) (*JWTClaims, error) {
	//解析声明  传一个token 和声明 一个方法
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify")
	return claims, nil
}

//func verifyAction(strToken string) (*JWTClaims, error) {
//	//解析声明  传一个token 和声明 一个方法
//	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(Secret), nil
//	})
//	if err != nil {
//		return nil, errors.New(ErrorReason_ServerBusy)
//	}
//	claims, ok := token.Claims.(*JWTClaims)
//	if !ok {
//		return nil, errors.New(ErrorReason_ReLogin)
//	}
//	if err := token.Claims.Valid(); err != nil {
//		return nil, errors.New(ErrorReason_ReLogin)
//	}
//	fmt.Println("verify")
//	return claims, nil
//}
func getToken(claims *JWTClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "",errors.New(ErrorReason_ServerBusy)
	}
	return signedToken,nil
}