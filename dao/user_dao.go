package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/capyflow/Allspark-go/ds"
	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/model"
	"github.com/redis/go-redis/v9"
)

// buildUserProfileKey 构建用户ID到用户详情的键
func buildUserProfileKey(group string) string {
	return fmt.Sprintf("%s:user:profile", group)
}

// UserDao 用户 DAO
type UserDao struct {
	ctx   context.Context
	group string
	rdb   *redis.Client
}

// NewNewUserDao 创建用户 DAO
func NewNewUserDao(ctx context.Context, group string, dServer *ds.DatabaseServer) *UserDao {
	rdb, ok := dServer.GetRedis("user")
	if !ok {
		panic("user redis not found")
	}
	return &UserDao{
		ctx:   ctx,
		group: group,
		rdb:   rdb,
	}
}

// QueryUserProfile 查询用户的详细信息
func (u *UserDao) QueryUserProfile(ctx context.Context) (*model.UserProfile, error) {
	key := buildUserProfileKey(u.group)
	raw, err := u.rdb.Get(ctx, key).Result()
	if err != nil {
		logx.Errorf("UserDao|QueryUserProfile|Error|%v|%s", err, key)
		return nil, err
	}
	profile := &model.UserProfile{}
	err = json.Unmarshal([]byte(raw), profile)
	if nil != err {
		logx.Errorf("UserDao|QueryUserProfile|Error|%v|%s", err, key)
		return nil, err
	}
	return profile, nil
}

// 修改用户信息
func (u *UserDao) UpdateUserProfile(ctx context.Context, profile *model.UserProfile) error {
	key := buildUserProfileKey(u.group)
	profile.UpdateTs = time.Now().Unix()
	raw, err := json.Marshal(profile)
	if nil != err {
		logx.Errorf("UserDao|UpdateUserProfile|Error|%v|%s", err, key)
		return err
	}
	err = u.rdb.Set(ctx, key, raw, 0).Err()
	if err != nil {
		logx.Errorf("UserDao|UpdateUserProfile|Error|%v|%s", err, key)
		return err
	}
	return nil
}
