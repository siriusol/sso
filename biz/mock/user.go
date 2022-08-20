package mock

import (
	"context"
	"github.com/siriusol/sso/biz/model"
)

func GetUserInfo(ctx context.Context, userId int64) (*model.User, error) {
	switch userId {
	case 142857:
		return &model.User{
			UserId: 142857,
			AccountIds: []string{
				"0001_1",
				"0018_1",
				"0018_2",
				"0018_3",
				"0023_1",
			},
		}, nil
	}
	return nil, nil
}
