package sso

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siriusol/sso/biz/mock"
	"github.com/siriusol/sso/biz/model"
	"github.com/siriusol/sso/biz/pkg/cookie"
	"github.com/siriusol/sso/biz/pkg/session"
	"github.com/siriusol/sso/biz/pkg/session_lib"
)

func Login(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId64, _ := strconv.ParseInt(userIdStr, 10, 64)
	accountTypeStr := c.Query("account_type")
	accountTypeI64, _ := strconv.ParseInt(accountTypeStr, 10, 64)
	appIdStr := c.Query("app_id")
	appIdI64, _ := strconv.ParseInt(appIdStr, 10, 64)
	accountId := c.Query("account_id")
	if userId64 == 0 || accountTypeI64 == 0 || appIdI64 == 0 || accountId == "" {
		log.Printf("[Login] invalid request, userId=%s, accountType=%s, appId=%s, accoundId=%s", userIdStr, accountTypeStr, appIdStr, accountId)
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "invalid request",
		})
		return
	}
	ctx := c.Copy()

	var (
		sm          model.SessionModel
		accountType = int32(accountTypeI64)
		appId       = int32(appIdI64)
	)

	ck, err := c.Cookie("session_key")
	if err != nil || ck == "" {
		// first login
		sm = newSession(ctx, userId64, accountType, appId, accountId)
	} else {
		err = cookie.Unmarshal([]byte(ck), &sm)
		if err != nil {
			// warn
			log.Printf("[UserInfo] session key in cookie is invalid")
			sm = newSession(ctx, userId64, accountType, appId, accountId)
		} else {
			var remoteSM *model.SessionModel
			remoteSM, err = mock.GetSession(ctx, sm.SessionKey)
			if err != nil {
				// error
				log.Printf("[UserInfo.GetSession] error=%+v", err)
				c.JSON(http.StatusOK, map[string]interface{}{
					"msg": "please try later",
				})
				return
			}
			if session.DiffSM(&sm, remoteSM) {
				// warn
				sm = newSession(ctx, userId64, accountType, appId, accountId)
			} else {
				updateSession(ctx, &sm, accountType, appId, accountId)
			}
		}
	}

	user, _ := mock.GetUserInfo(ctx, userId64)
	if user == nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "user not exist",
		})
		return
	}

	v, _ := cookie.Marshal(sm)
	c.SetCookie("session_key", string(v), 600, "/", c.Request.Host, false, false)
	c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "success",
	})
	return
}

func newSession(ctx context.Context, userId int64, accountType int32, appId int32, accountId string) model.SessionModel {
	n := time.Now()
	sm := model.SessionModel{
		UserId:    userId,
		AccountId: accountId,
		LoginMap: map[int32]map[int32]string{
			accountType: {
				appId: accountId,
			},
		},
		CreatedTime: n.Unix(),
		UpdatedTime: n.Unix(),
		TTL:         600, // 600s / 10 min
	}
	sessionKey := session_lib.NewSimpleSessionKeyMaker().Make(&sm)
	sm.SessionKey = sessionKey
	_ = mock.AddSession(ctx, sessionKey, &sm)
	return sm
}

func updateSession(ctx context.Context, sm *model.SessionModel, accountType int32, appId int32, accountId string) {
	if sm == nil {
		return
	}
	if sm.LoginMap == nil {
		newSession(ctx, sm.UserId, sm.AccountType, sm.AppId, sm.AccountId)
		return
	}
	n := time.Now()
	sm.AccountId = accountId
	if sm.LoginMap[accountType] == nil {
		sm.LoginMap[accountType] = make(map[int32]string)
	}
	sm.LoginMap[accountType][appId] = accountId
	sm.UpdatedTime = n.Unix()
	_ = mock.UpdateSession(ctx, sm.SessionKey, sm)
}
