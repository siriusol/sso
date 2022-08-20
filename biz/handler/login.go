package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siriusol/sso/biz/mock"
	"github.com/siriusol/sso/biz/model"
	"github.com/siriusol/sso/biz/pkg/cookie"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId64, _ := strconv.ParseInt(userIdStr, 10, 64)
	accountType := c.Query("account_type")
	accountId := c.Query("account_id")
	if userId64 == 0 || accountType == "" || accountId == "" {
		log.Printf("[Login] invalid serId or accountId, userId=%s, accound_id=%s", userIdStr, accountId)
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "invalid request",
		})
		return
	}
	ctx := c.Copy()

	var sm model.SessionModel
	ck, err := c.Cookie("session_key")
	if err != nil || ck == "" {
		// first login
		sm = newSession(ctx, userId64, accountType, accountId)
	} else {
		err = cookie.Unmarshal([]byte(ck), &sm)
		if err != nil {
			// warn
			log.Printf("[UserInfo] session key in cookie is invalid")
			sm = newSession(ctx, userId64, accountType, accountId)
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
			if diffSM(&sm, remoteSM) {
				// warn
				sm = newSession(ctx, userId64, accountType, accountId)
			} else {
				updateSession(ctx, &sm, accountType, accountId)
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

func newSession(ctx context.Context, userId int64, accountType string, accountId string) model.SessionModel {
	n := time.Now()
	sessionKey := fmt.Sprint(userId)
	sm := model.SessionModel{
		UserId:    userId,
		AccountId: accountId,
		LoginMap: map[string]string{
			accountType: accountId,
		},
		SessionKey:  sessionKey,
		CreatedTime: n.Unix(),
		UpdatedTime: n.Unix(),
		TTL:         600, // 600s / 10 min
	}
	_ = mock.AddSession(ctx, sessionKey, &sm)
	return sm
}

func updateSession(ctx context.Context, sm *model.SessionModel, accountType string, accountId string) {
	n := time.Now()
	sm.AccountId = accountId
	sm.LoginMap[accountType] = accountId
	sm.UpdatedTime = n.Unix()
	_ = mock.UpdateSession(ctx, sm.SessionKey, sm)
}

func diffSM(a, b *model.SessionModel) bool {
	if a == nil || b == nil {
		return true
	}
	if a.UserId != b.UserId || a.AccountId != b.AccountId || a.SessionKey != b.SessionKey || a.CreatedTime != b.CreatedTime || a.UpdatedTime != b.UpdatedTime {
		return true
	}
	//if len(a.LoginMap) != len(b.LoginMap) || len(a.LoginMap) == 0 {
	//	return true
	//}
	//for k, av := range a.LoginMap {
	//	if bv, ok := b.LoginMap[k]; !ok {
	//		return true
	//	} else if av != bv {
	//		return true
	//	}
	//}
	return false
}
