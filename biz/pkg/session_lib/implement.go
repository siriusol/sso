package session_lib

import (
	"fmt"
	"github.com/siriusol/sso/biz/model"
)

type SimpleSessionKeyMaker struct{}

func NewSimpleSessionKeyMaker() *SimpleSessionKeyMaker {
	return &SimpleSessionKeyMaker{}
}

func (m *SimpleSessionKeyMaker) Make(sm *model.SessionModel) string {
	if sm == nil {
		return ""
	}
	return fmt.Sprintf("session_key_%d_%d_%d_%s", sm.UserId, sm.AccountType, sm.AppId, sm.AccountId)
}
