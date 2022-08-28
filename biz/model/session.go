package model

/**
业务线 -> 端 -> 账户
accountType -> appId -> accountId
*/

type SessionModel struct {
	UserId      int64                      `json:"user_id"`
	AccountType int32                      `json:"account_type"` // 兼容历史，当前登录的业务线
	AppId       int32                      `json:"app_id"`       // 兼容历史，当前登录的端 ID
	AccountId   string                     `json:"account_id"`   // 兼容历史，当前登录的账户
	LoginMap    map[int32]map[int32]string `json:"login_map"`    // 已登录的账户，限制每个业务线下每个端最多同时登录一个账户 map[accountType]map[appId]accountId
	SessionKey  string                     `json:"session_key"`  // session key
	CreatedTime int64                      `json:"created_time"` // 创建时间
	UpdatedTime int64                      `json:"updated_time"` // 更新时间
	TTL         int64                      `json:"ttl"`          // session 有效时间段，单位：秒
}

type User struct {
	UserId   int64
	Accounts []*Account // 相同账号，可在不同业务线下有多个账户，也可在同一业务线下有多个账户
}

type Account struct {
	AccountType int32
	AccountId   string
}
