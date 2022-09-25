package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sparrow/internal/pkg/slog"
	. "sparrow/internal/pkg/util"
)

func NewImpl() *wechatImpl {
	return &wechatImpl{}
}

type wechatImpl struct {
}

func (*wechatImpl) getErrorInfo(body string) (*ErrorResp, error) {
	errresp := &ErrorResp{}
	if err := json.Unmarshal([]byte(body), errresp); err != nil {
		slog.Error("json unmarshal error: ", err)
		return nil, err
	}
	return errresp, nil
}

func (wi *wechatImpl) GetGrantInfo(ctx context.Context, req *GetGrantInfoReq) (*GrantInfo, error) {
	p := fmt.Sprintf("appid=%s&secret=%s&code=%s&grant_type=%s", req.AppId, req.AppSecret, req.Code, req.GrantType)
	_, body, err := HttpGet(ctx, "https://api.weixin.qq.com/sns/oauth2/access_token", WithParams(p))
	if err != nil {
		slog.Error("call wechat api error: ", err)
		return nil, err
	} else {
		grantInfo := &GrantInfo{}
		if err := json.Unmarshal([]byte(body), grantInfo); err == nil {
			return grantInfo, nil
		} else {
			errresp, err := wi.getErrorInfo(body)
			if err != nil {
				slog.Errorf("wechat body error: %s  body: %s", err.Error(), body)
				return nil, err
			} else {
				text := fmt.Sprintf("call wechat api response error. errorcode: %d  errormsg: %s", errresp.ErrCode, errresp.ErrMsg)
				slog.Errorf(text)
				return nil, errors.New(text)
			}
		}
	}
}

func (wi *wechatImpl) GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*UserInfo, error) {
	p := fmt.Sprintf("access_token=%s&openid=%s", req.AccessToken, req.OpenId)
	_, body, err := HttpGet(ctx, "https://api.weixin.qq.com/sns/userinfo", WithParams(p))
	if err != nil {
		slog.Error("call wechat api error: ", err)
		return nil, err
	} else {
		userInfo := &UserInfo{}
		if err := json.Unmarshal([]byte(body), userInfo); err == nil {
			return userInfo, nil
		} else {
			errresp, err := wi.getErrorInfo(body)
			if err != nil {
				slog.Errorf("wechat body error: %s  body: %s", err.Error(), body)
				return nil, err
			} else {
				text := fmt.Sprintf("call wechat api response error. errorcode: %d  errormsg: %s", errresp.ErrCode, errresp.ErrMsg)
				slog.Errorf(text)
				return nil, errors.New(text)
			}
		}
	}
}
