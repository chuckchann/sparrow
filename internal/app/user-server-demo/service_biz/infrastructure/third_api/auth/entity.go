package wechat

type ErrorResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type (
	GetGrantInfoReq struct {
		AppId     string `json:"app_id"`
		AppSecret string `json:"secret"`
		Code      string `json:"code"`
		GrantType string `json:"grant_type"`
	}

	GrantInfo struct {
		AccessToken  string `json:"access_token"`
		ExpireIn     int    `json:"expire_in"`
		RefreshToken string `json:"refresh_token"`
		OpenId       string `json:"openid"`
		Scope        string `json:"scope"`
		UnionId      string `json:"unionid"`
	}
)

type (
	GetUserInfoReq struct {
		AccessToken string `json:"access_token"`
		OpenId      string `json:"openid"`
	}

	UserInfo struct {
		OpenId     string   `json:"openid"`
		NickName   string   `json:"nickname"`
		Sex        int      `json:"sex"`
		Province   string   `json:"province"`
		City       string   `json:"city"`
		Country    string   `json:"country"`
		HeadImgUrl string   `json:"headimgurl"`
		Privilege  []string `json:"privilege"`
		UnionId    string   `json:"unionid"`
	}
)

type (
	RefreshTokenReq struct {
		AppId        string `json:"app_id"`
		GrantType    string `json:"grant_type"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenResp struct {
		AccessToken  string `json:"access_token"`
		ExpireIn     int    `json:"expire_in"`
		RefreshToken string `json:"refresh_token"`
		OpenId       string `json:"openid"`
		Scope        string `json:"scope"`
	}
)
