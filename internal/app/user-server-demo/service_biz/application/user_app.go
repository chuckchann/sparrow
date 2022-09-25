package application

import (
	"context"
	"sparrow/internal/app/user-server-demo/service_biz/domain/entity"
	"sparrow/internal/app/user-server-demo/service_biz/domain/service"
	. "sparrow/internal/pkg/entity"
	"sparrow/internal/pkg/serror"
)

//应用接入层 类似Controller层
//TODO: use go-validate to validate request param

type UserAppInterface interface {
	SaveUser(context.Context, *entity.User) *Response
	GetUsers(context.Context) ([]entity.User, *Response)
	GetUser(context.Context, uint64) (*entity.User, *Response)
	Login(context.Context, string, string) *Response
}

//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

type userApp struct {
	us *service.UserService //business logic
}

func (u *userApp) SaveUser(ctx context.Context, req *entity.User) *Response {
	errcode := u.us.SaveUser(req)
	if errcode != nil {
		return errcode.Response()
	}
	return nil
}

func (u *userApp) GetUser(ctx context.Context, userId uint64) (*entity.User, *Response) {
	if userId <= 0 {
		return nil, serror.ERR_PARAM.Response()
	}
	user, errcode := u.us.GetUser(userId)
	if errcode != nil {
		return nil, errcode.Response()
	}
	return user, nil
}

func (u *userApp) GetUsers(ctx context.Context) ([]entity.User, *Response) {
	user, errcode := u.us.GetUsers()
	if errcode != nil {
		return nil, errcode.Response()
	}
	return user, nil
}

func (u *userApp) Login(ctx context.Context, code, grantType string) *Response {
	token, errcode := u.us.Login(ctx, code, grantType)
	if errcode != nil {
		return errcode.Response()
	}
	return serror.SUCCESS.Response().ReplaceData(token)
}
