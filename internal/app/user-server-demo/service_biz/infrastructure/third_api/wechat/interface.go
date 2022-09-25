package wechat

import "context"

var _ Interface = (*wechatImpl)(nil)

var api Interface = NewImpl()

func SetApi(a Interface) {
	api = a
}

type Interface interface {
	GetGrantInfo(context.Context, *GetGrantInfoReq) (*GrantInfo, error)
	GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*UserInfo, error)
}

func GetGrantInfo(ctx context.Context, req *GetGrantInfoReq) (*GrantInfo, error) {
	return api.GetGrantInfo(ctx, req)
}

func GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*UserInfo, error) {
	return api.GetUserInfo(ctx, req)
}

func RefreshToken(ctx context.Context, req *RefreshTokenReq) {

}
