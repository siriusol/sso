package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siriusol/sso/biz/handler/business"
	"github.com/siriusol/sso/biz/handler/sso"
	"github.com/siriusol/sso/biz/pkg/session_lib"
)

func main() {
	e := gin.Default()
	e.POST("/sso/login", sso.Login)
	e.POST("/sso/logout", sso.Logout)
	e.POST("/sso/register", sso.Register)
	e.POST("/sso/user_info", sso.UserInfo)

	businessAGroup := e.Group("/a")
	businessAGroup.Use(session_lib.Middleware)
	businessAGroup.GET("/user_info", business.UserInfo)

	businessBGroup := e.Group("/b")
	businessBGroup.Use(session_lib.Middleware)
	businessBGroup.GET("/user_info", business.UserInfo)

	businessCGroup := e.Group("/c")
	businessCGroup.Use(session_lib.Middleware)
	businessCGroup.GET("/user_info", business.UserInfo)

	e.Run(":8080")
}
