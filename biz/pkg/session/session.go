package session

import "github.com/siriusol/sso/biz/model"

func DiffSM(a, b *model.SessionModel) bool {
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
