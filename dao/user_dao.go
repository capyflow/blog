package dao

import (
	"context"
	"encoding/json"

	"github.com/capyflow/Allspark-go/ds"
	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/model"
	"github.com/redis/go-redis/v9"
)

func buildUserProfileKey(email string) string {
	return "user:profile:" + email
}

// UserDao 用户 DAO
type UserDao struct {
	ctx context.Context
	rdb *redis.Client
}

// NewNewUserDao 创建用户 DAO
func NewNewUserDao(ctx context.Context, dServer *ds.DatabaseServer) *UserDao {
	rdb, ok := dServer.GetRedis("user")
	if !ok {
		panic("user redis not found")
	}
	return &UserDao{
		ctx: ctx,
		rdb: rdb,
	}
}

// QueryUserProfile 查询用户的详细信息
func (u *UserDao) QueryUserProfile(ctx context.Context, email string) (*model.UserProfile, error) {
	key := buildUserProfileKey(email)
	profile := &model.UserProfile{}
	err := u.rdb.Get(ctx, key).Scan(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// SaveUserProfile 保存用户的详细信息
func (u *UserDao) SaveUserProfile(ctx context.Context, email string, profile *model.UserProfile) error {
	key := buildUserProfileKey(email)
	jsonData, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	err = u.rdb.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		logx.Errorf("UserDao|SaveUserProfile|Error|%v|%s", err, key)
		return err
	}
	return nil
}
