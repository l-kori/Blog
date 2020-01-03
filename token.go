package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/login/:username/:password", login)
	r.GET("/verify/:token", verify)
	r.GET("/refresh/:token", refresh)
	r.GET("/sayHello/:token", sayHello)
	r.Run(":9090")

	//http://localhost:9090/login/dong/123456
	//http://localhost:9090/verify/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA1MTIyMTAsImlhdCI6MTU2MDUwODYxMCwidXNlcl9pZCI6MSwicGFzc3dvcmQiOiIxMjM0NTYiLCJ1c2VybmFtZSI6ImRvbmciLCJmdWxsX25hbWUiOiJkb25nIiwicGVybWlzc2lvbnMiOltdfQ.Esh1Zge0vO1BAW1GeR5wurWP3H1jUIaMf3tcSaUwkzA
	//http://localhost:9090/refresh/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA1MTIyNDMsImlhdCI6MTU2MDUwODYxMCwidXNlcl9pZCI6MSwicGFzc3dvcmQiOiIxMjM0NTYiLCJ1c2VybmFtZSI6ImRvbmciLCJmdWxsX25hbWUiOiJkb25nIiwicGVybWlzc2lvbnMiOltdfQ.Xkb_J8MWXkwGUcBF9bpp2Ccxp8nFPtRzFzOBeboHmg0
}

const(
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin = "请重新登陆"
)

func sayHello(c *gin.Context) {
	strToken := c.Param("token")
	claim,err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "hello,",claim.Username)
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

func login(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
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
	c.String(http.StatusOK, signedToken)
}

func verify(c *gin.Context) {
	strToken := c.Param("token")
	claim,err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "verify,",claim.Username)
}

func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims,err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken,err:=getToken(claims)
	if err!=nil{
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

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

func getToken(claims *JWTClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "",errors.New(ErrorReason_ServerBusy)
	}
	return signedToken,nil
}