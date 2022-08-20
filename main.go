package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siriusol/sso/biz/handler"
)

func main() {
	e := gin.Default()
	e.POST("/sso/login", handler.Login)
	e.POST("/sso/logout", handler.Logout)
	e.POST("/sso/register", handler.Register)
	e.POST("/sso/user_info", handler.UserInfo)

	e.Run(":8080")
}
