package mock

import (
	"context"
	"errors"
	"github.com/siriusol/sso/biz/model"
)

/**
Session 服务的 Mock 实现，用于根据 session key 操作跨业务线的 session
*/

var global = map[string]*model.SessionModel{
	"9_9999_999999_9999_999": &model.SessionModel{
		UserId:      9,
		AccountType: 9999,
		AppId:       999999,
		AccountId:   "9999_999",
		LoginMap: map[int32]map[int32]string{
			9999: {
				999999: "9999_99",
			},
		},
		SessionKey: "9_9999_999999_9999_999",
	},
}

var Err_SessionNotExist = errors.New("session not exist")

func GetSession(ctx context.Context, sessionKey string) (*model.SessionModel, error) {
	return global[sessionKey], nil
}

func AddSession(ctx context.Context, sessionKey string, sm *model.SessionModel) error {
	global[sessionKey] = sm
	return nil
}

func UpdateSession(ctx context.Context, sessionKey string, sm *model.SessionModel) error {
	if _, ok := global[sessionKey]; !ok {
		return Err_SessionNotExist
	}
	global[sessionKey] = sm
	return nil
}
