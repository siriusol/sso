package mock

import (
	"context"
	"github.com/siriusol/sso/biz/model"
)

/**
账号服务的 Mock 实现
*/

func GetUserInfo(ctx context.Context, userId int64) (*model.User, error) {
	switch userId {
	case 142857:
		return &model.User{
			UserId: 142857,
			Accounts: []*model.Account{
				{
					1, "0001_1",
				},
				{
					18, "0018_1",
				},
				{
					18, "2",
				},
				{
					18, "3",
				},
				{
					23, "1",
				},
			},
		}, nil
	}
	return nil, nil
}
