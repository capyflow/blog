package service

import (
	"context"
	"errors"

	"github.com/capyflow/Allspark-go/jwtx"
	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/api"
	"github.com/capyflow/blog/dao"
	"github.com/capyflow/blog/model"
	"github.com/capyflow/blog/pkg"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx         context.Context
	userDao     *dao.UserDao
	secret      string
	uid         string
	userAccount struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

// NewUserService 创建用户服务
func NewUserService(ctx context.Context, cfg *pkg.Config, userDao *dao.UserDao) *UserService {
	us := &UserService{
		ctx:     ctx,
		secret:  cfg.Server.ConsoleJwt.Secret,
		userDao: userDao,
		userAccount: struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{
			Email:    *cfg.BlogUser.Email,
			Password: EncryptPassword(*cfg.BlogUser.Password),
		},
		uid: uuid.NewString(),
	}
	us.initUserProfile()
	return us
}

// 初始化用户信息
func (u *UserService) initUserProfile() error {
	_, err := u.QueryUserProfile(u.ctx)
	if nil != err && !errors.Is(err, redis.Nil) {
		logx.Errorf("UserService|initUserProfile|QueryUserProfile failed|%v", err)
		return err
	}
	if errors.Is(err, redis.Nil) {
		if err := u.userDao.UpdateUserProfile(u.ctx, &model.UserProfile{Nickname: "Aaron", Avatar: ""}); nil != err {
			logx.Errorf("UserService|initUserProfile|UpdateUserProfile failed|%v", err)
			return err
		}
	}
	return nil
}

// login 登录
func (u *UserService) LoginByPwd(ctx context.Context, email, password string) (string, error) {
	if u.userAccount.Email != email {
		return "", pkg.ErrorsEnum.ErrEmailNotMatch
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.userAccount.Password), []byte(password)); err != nil {
		return "", pkg.ErrorsEnum.ErrPasswordNotMatch
	}
	// 密码匹配，返回用户信息，生成token
	token, err := jwtx.SignJwt(u.secret, jwt.MapClaims{
		"uid": u.uid,
	})
	if err != nil {
		logx.Errorf("UserService|LoginByPwd|Error|%v|%s", err, email)
		return "", err
	}
	return token, nil
}

// 根据验证码登录
func (u *UserService) LoginByCode(ctx context.Context, email, code string) (string, error) {
	return "", errors.New("not implemented")
}

// 密码加密
func EncryptPassword(pwd string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		logx.Errorf("UserService|EncryptPassword|Error|%v|%s", err, pwd)
		return ""
	}
	return string(hashed)
}

// queryUserProfile 查询用户信息
func (u *UserService) QueryUserProfile(ctx context.Context) (*model.UserProfile, error) {
	return u.userDao.QueryUserProfile(ctx)
}

// 更新用户信息
func (u *UserService) UpdateUserProfile(ctx context.Context, updateInfo *api.UpdateUserProfileReq) (*model.UserProfile, error) {
	userProfile, err := u.QueryUserProfile(ctx)
	if nil != err {
		logx.Errorf("UserService|UpdateUserProfile|QueryUserProfile failed|%v", err)
		return nil, err
	}
	if len(updateInfo.Avatar) > 0 {
		userProfile.Avatar = updateInfo.Avatar
	}
	if len(updateInfo.Nickname) > 0 {
		userProfile.Nickname = updateInfo.Nickname
	}
	return userProfile, u.userDao.UpdateUserProfile(ctx, userProfile)
}
