package session_lib

import (
	"github.com/gin-gonic/gin"
	"github.com/siriusol/sso/biz/model"
)

func Middleware(c *gin.Context) {
	preRequest(c)
	c.Next()
	preResponse(c)
}

func preRequest(c *gin.Context) {
	// parse cookie
	// get session
	// set user info in context
}

func GetUserId() int64 {
	return getSessionModel().UserId
}

func GetAccountType() int32 {
	return getSessionModel().AccountType
}

func GetAppId() int32 {
	return getSessionModel().AppId
}

func GetAccountId() string {
	return getSessionModel().AccountId
}

func getSessionModel() *model.SessionModel {
	return &model.SessionModel{}
}

func preResponse(c *gin.Context) {

}
