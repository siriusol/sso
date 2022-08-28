package session_lib

import (
	"github.com/siriusol/sso/biz/model"
)

type SessionKeyMaker interface {
	Make(m *model.SessionModel) string
}
