package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sparrow/internal/app/user-server-demo/service_biz/domain/entity"
	"sparrow/internal/app/user-server-demo/service_biz/domain/repo"
	"sparrow/internal/app/user-server-demo/service_biz/infrastructure/persistence"
	"sparrow/internal/app/user-server-demo/service_biz/infrastructure/third_api/wechat"
	. "sparrow/internal/pkg/entity"
	"sparrow/internal/pkg/serror"
	"time"
)

//service layer is for domain logic

type UserService struct {
	userRepo persistence.UserRepo
}

func (us *UserService) SaveUser(u *entity.User) *serror.SError {
	_, err := us.userRepo.SaveUser(u)
	if err != nil {
		return serror.ERR_SYSTEM
	}
	return nil
}

func (us *UserService) GetUsers() ([]entity.User, *serror.SError) {
	users, err := us.userRepo.GetUsers()
	if err != nil {
		return nil, serror.ERR_SYSTEM
	}
	return users, nil
}

func (us *UserService) GetUser(id uint64) (*entity.User, *serror.SError) {
	user, err := us.userRepo.GetUser(id)
	if err != nil {
		return nil, serror.ERR_SYSTEM
	}
	return user, nil
}

func (us *UserService) Login(ctx context.Context, code, grantType string) (string, *serror.SError) {
	//get grant info(access_token) by code
	var err error
	grantInfo, err := wechat.GetGrantInfo(ctx, &wechat.GetGrantInfoReq{
		AppId:     viper.GetString("wechat.appId"),
		AppSecret: viper.GetString("wecaht.appSecret"),
		Code:      code,
		GrantType: grantType,
	})
	if err != nil {
		return "", serror.ERR_CALL
	}

	//get user info by access token
	wechatUser, err := wechat.GetUserInfo(ctx, &wechat.GetUserInfoReq{
		AccessToken: grantInfo.AccessToken,
		OpenId:      grantInfo.OpenId,
	})
	if err != nil {
		return "", serror.ERR_CALL
	}

	//check user
	var user *entity.User
	user, err = us.userRepo.GetUserByUnionId(wechatUser.UnionId)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return "", serror.ERR_SYSTEM
	} else if gorm.IsRecordNotFoundError(err) {
		//register user
		user, err = us.userRepo.SaveUser(&entity.User{
			FirstName: wechatUser.NickName,
			Sex:       wechatUser.Sex,
			UnionId:   wechatUser.UnionId,
		})
		if err != nil {
			return "", serror.ERR_SYSTEM
		}
	}

	//generate jwt token
	jwtObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user.ID,
		"expire_time": time.Now().Add(time.Minute * time.Duration(viper.GetInt("token.ttl"))).Unix(),
	})
	token, err := jwtObj.SignedString(viper.GetString("token.secret"))
	if err != nil {
		return "", serror.ERR_SYSTEM
	}

	return token, nil
}
