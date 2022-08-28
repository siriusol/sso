package business

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siriusol/sso/biz/mock"
	"github.com/siriusol/sso/biz/model"
	"github.com/siriusol/sso/biz/pkg/cookie"
	"github.com/siriusol/sso/biz/pkg/notation"
	"github.com/siriusol/sso/biz/pkg/session"
)

func UserInfo(c *gin.Context) {
	ck, err := c.Cookie("session_key")
	if err != nil {
		log.Printf("[UserInfo] get cookie error=%+v", err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "user not login",
		})
		return
	}
	if ck == "" {
		log.Printf("[UserInfo] session key in cookie is empty")
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "user not login",
		})
		return
	}
	var sm model.SessionModel
	err = cookie.Unmarshal([]byte(ck), &sm)
	if err != nil {
		log.Printf("[UserInfo] session key in cookie is invalid")
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "user not login",
		})
		return
	}
	ctx := c.Copy()
	remoteSM, _ := mock.GetSession(ctx, sm.SessionKey)
	if session.DiffSM(&sm, remoteSM) {
		log.Printf("[UserInfo] session diff")
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "user not login",
		})
		return
	}
	log.Printf("[UserInfo] user=%s", notation.JSONStr(sm))
	c.JSON(http.StatusOK, sm)
}
