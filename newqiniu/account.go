package newqiniu

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gohttp"
)

//------------------获取账户管理凭证

//获取账户管理凭证返回数据结构
type AccountOauthTokenResponse struct {
	QiniuBaseResponse
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// 获取账户管理凭证请求数据结构
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

//Accesstoken数据结构
type Access struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"-"`
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Uid          int64  `json:"uid"`
	Email        string `json:"email"`
	Userid       string `json:"userid"`
	Key          string `json:"key"`
	Secret       string `json:"secret"`
}

//获取Access缓存文件名
func (access *Access) GetFileName() string {
	if len(access.Username) == 0 {
		panic("username not set")
	}
	return access.Username + ".json"
}

//是否过期
func (access *Access) IsExpired() bool {
	return access.ExpiresIn <= (time.Now().Unix() - 60)
}

//获取Accesstoken
func (access *Access) Get() error {
	if len(access.Username) == 0 || len(access.Password) == 0 {
		return fmt.Errorf("missing account username and password")
	}
	req := &AccountOauthTokenRequest{}
	req.Username = access.Username
	req.Password = access.Password
	resp := &AccountOauthTokenResponse{}
	err := gohttp.DoPostResponse(req, resp)
	if err != nil {
		return err
	} else {
		if resp.OK() {
			access.AccessToken = resp.AccessToken
			access.ExpiresIn = time.Now().Unix() + resp.ExpiresIn
			access.RefreshToken = resp.RefreshToken

			// 获取账号信息
			ireq := &AccountInfoRequest{}
			ireq.AuthorizationValue = access.AccessToken
			iresp := &AccountInfoResponse{}
			err := gohttp.DoPostResponse(ireq, iresp)
			if err != nil {
				return err
			} else {
				if iresp.OK() {
					access.Uid = iresp.UID
					access.Userid = iresp.Userid
					access.Email = iresp.Email
				} else {
					return fmt.Errorf("%v", iresp.Msg())
				}
			}

			// 获取账号AK/SK
			kreq := &AccountKeysRequest{
				App: "default",
			}
			kreq.AuthorizationValue = access.AccessToken
			kresp := &AccountKeysResponse{}
			err2 := gohttp.DoPostResponse(kreq, kresp)
			if err2 != nil {
				return err2
			} else {
				if kresp.OK() {
					access.Key = kresp.Key
					access.Secret = kresp.Secret
				} else {
					return fmt.Errorf("%v", kresp.Msg())
				}
			}

			return access.WriteOauth()
		} else {
			return fmt.Errorf("%v", resp.Msg())
		}
	}
}

//刷新Accesstoken
func (access *Access) Refresh(force bool) error {
	if force || access.IsExpired() {
		req := AccountRefreshOauthTokenRequest{}
		req.RefreshToken = access.RefreshToken
		resp := &AccountRefreshOauthTokenResponse{}
		err := gohttp.DoPostResponse(&req, resp)
		if err != nil {
			return err
		} else {
			if resp.OK() {
				access.AccessToken = resp.AccessToken
				access.ExpiresIn = time.Now().Unix() + resp.ExpiresIn
				access.RefreshToken = resp.RefreshToken
				return access.WriteOauth()
			} else {
				return fmt.Errorf("%v", resp.Msg())
			}
		}
	}
	return nil
}

//将accesstoken写入文件
func (access *Access) WriteOauth() error {
	if bytes, err := json.Marshal(access); err != nil {
		return err
	} else {
		if accesspath := GetDir(); len(accesspath) == 0 {
			return err
		} else {
			err := os.MkdirAll(accesspath, 0777)
			if err != nil {
				return err
			}
			if err := gofile.Json2File(string(bytes), accesspath+string(os.PathSeparator)+access.GetFileName(), 0777); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}

//从文件读取accesstoken
func (access *Access) ReadOauth(autorefresh bool) error {
	accesspath := GetDir() + string(os.PathSeparator) + access.GetFileName()
	if err := gofile.File2Json(accesspath, &access); err != nil {
		return err
	} else {
		if access.IsExpired() && autorefresh {
			err := access.Refresh(true)
			if err != nil {
				return access.Get()
			}
		}
		return nil
	}
}
