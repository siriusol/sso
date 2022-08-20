package mock

import (
	"context"
	"errors"
	"github.com/siriusol/sso/biz/model"
)

var global = map[string]*model.SessionModel{
	"ping": &model.SessionModel{
		UserId:    9999,
		AccountId: "9999_99",
		LoginMap: map[string]string{
			"9999": "9999_99",
		},
		SessionKey: "ping",
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
