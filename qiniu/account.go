package qiniu

import (
	"time"
)

//------------------获取账户管理凭证

//获取账户管理凭证返回数据结构
type AccountOauthTokenResponse struct {
	QiniuBaseResponse
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

//获取账户管理凭证请求数据结构
type AccountOauthTokenRequest struct {
	QiniuBaseRequest
	GrantType string `url:"grant_type"`
	Username  string `url:"username"`
	Password  string `url:"password"`
}

//填充默认信息
func (req *AccountOauthTokenRequest) B() error {
	if err := req.QiniuBaseRequest.B(); err != nil {
		return err
	}
	req.host = HOST_ACC
	req.GrantType = "password"
	return nil
}

//获取账户管理凭证请求链接
func (req *AccountOauthTokenRequest) URL() string {
	return "https://" + HOST_ACC + "/oauth2/token"
}

//------------------刷新账户管理凭证

//刷新账户管理凭证返回数据结构
type AccountRefreshOauthTokenResponse struct {
	QiniuBaseResponse
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

//刷新账户管理凭证请求数据结构
type AccountRefreshOauthTokenRequest struct {
	QiniuBaseRequest
	GrantType    string `url:"grant_type"`
	RefreshToken string `url:"refresh_token"`
}

//刷新账户管理凭证请求链接
func (req *AccountRefreshOauthTokenRequest) URL() string {
	return "https://" + HOST_ACC + "/oauth2/token"
}

//填充默认信息
func (req *AccountRefreshOauthTokenRequest) B() error {
	if err := req.QiniuBaseRequest.B(); err != nil {
		return err
	}
	req.host = HOST_ACC
	req.GrantType = "refresh_token"
	return nil
}

//------------------获取账户信息

//获取账户信息返回数据结构
type AccountInfoResponse struct {
	QiniuBaseResponse
	CanGetChildKey bool   `json:"can_get_child_key,omitempty"`
	DeviceNum      int    `json:"device_num,omitempty"`
	Email          string `json:"email"`
	InvitationNum  int    `json:"invitation_num,omitempty"`
	IsActivatied   bool   `json:"is_activated,omitempty"`
	IsDisabled     bool   `json:"is_disabled,omitempty"`
	ParentUID      int    `json:"parent_uid"`
	UID            int64  `json:"uid"`
	UserType       int    `json:"user_type,omitempty"`
	Userid         string `json:"userid"`
}

//获取账户信息请求数据结构
type AccountInfoRequest struct {
	QiniuBaseRequest
}

//获取账户信息请求链接
func (req *AccountInfoRequest) URL() string {
	return "https://" + HOST_ACC + "/user/info"
}

//填充默认信息
func (req *AccountInfoRequest) B() error {
	req.host = HOST_ACC
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//----------------获取主账号AK/SK

//获取主账号AK/SK返回数据结构
type AccountKeysResponse struct {
	QiniuBaseResponse
	Key           string    `json:"key"`
	Key2          string    `json:"key2"`
	Secret        string    `json:"secret"`
	Secret2       string    `json:"secret2"`
	State         int64     `json:"state"`
	State2        int64     `json:"state2"`
	AppID         int64     `json:"appId"`
	AppName       string    `json:"appName"`
	CreationTime  time.Time `json:"creation-time"`
	CreationTime2 time.Time `json:"creation-time2"`
	LastModified  time.Time `json:"last-modified"`
	LastModified2 time.Time `json:"last-modified2"`
	Uid           int64     `json:"uid"`
}

//获取主账号AK/SK请求数据结构
type AccountKeysRequest struct {
	QiniuBaseRequest
	App string `url:"app"`
}

//获取主账号AK/SK请求地址
func (req *AccountKeysRequest) URL() string {
	return "http://" + HOST_UC + "/appInfo"
}

//填充默认信息
func (req *AccountKeysRequest) B() error {
	req.host = HOST_UC
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

type AccountStatInfoResponse struct {
	QiniuBaseResponse
	ApicallGet int64 `json:"apicall_get"`
	ApicallPut int64 `json:"apicall_put"`
	Bandwidth  int64 `json:"bandwidth"`
	OvTransfer int64 `json:"ov_transfer"`
	Space      int64 `json:"space"`
	SpaceAvg   int64 `json:"space_avg"`
	Transfer   int64 `json:"transfer"`
}

type AccountStatInfoRequest struct {
	QiniuBaseRequest
	Uid    int64  `url:"uid,omitempty"`
	Bucket string `url:"bucket,omitempty"`
	Month  string `url:"month,omitempty"`
}

func (req *AccountStatInfoRequest) URL() string {
	return "http://" + HOST_API + "/stat/info"
}

func (req *AccountStatInfoRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}
