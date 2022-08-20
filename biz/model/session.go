package model

type SessionModel struct {
	UserId      int64             `json:"user_id"`
	AccountId   string            `json:"account_id"` // 兼容历史，当前登录的账户
	LoginMap    map[string]string `json:"login_map"`  // 已登录的业务线账户，限制每个业务线下最多同时登录一个账户
	SessionKey  string            `json:"session_key"`
	CreatedTime int64             `json:"created_time"`
	UpdatedTime int64             `json:"updated_time"`
	TTL         int64             `json:"ttl"` // session 有效时间段，单位：秒
}

type User struct {
	UserId     int64
	AccountIds []string
}
