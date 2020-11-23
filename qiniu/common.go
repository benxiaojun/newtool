package qiniu

import (
	"encoding/json"
	"os"
	"time"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gohttp"
	"net/url"
)

const HOST_ACC = "acc.qbox.me"
const HOST_UC = "uc.qbox.me"
const HOST_RS = "rs.qiniu.com"
const HOST_RSME = "rs.qbox.me"
const HOST_API = "api.qiniu.com"

const (
	BUCKET_PRIVATE = 0
	BUCKET_PUBLIC  = 1
)

//请求公共返回内容数据结构
type QiniuBaseResponse struct {
	Code             int    `json"code,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorCode        int    `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

//判断是否有错误
func (resp *QiniuBaseResponse) OK() bool {
	if len(resp.Error) == 0 {
		return true
	} else {
		return false
	}
}

//返回错误信息
func (resp *QiniuBaseResponse) Msg() string {
	if resp.Code > 0 {
		return fmt.Sprintf("[%d]%s", resp.Code, resp.Error)
	} else {
		return fmt.Sprintf("[%d]%s:%s", resp.ErrorCode, resp.Error, resp.ErrorDescription)
	}
}

const (
	AUTHORIZATION_BEARER = "Bearer" //主账号使用Authorization字段
	AUTHORIZATION_QBOX   = "QBox"   //子账号使用Authorization字段
)

//七牛请求公共参数数据结构
type QiniuBaseRequest struct {
	host               string `url:"-"` //请求头使用，不计入请求内容
	authorizationName  string `url:"-"` //请求头使用，不计入请求内容
	AuthorizationValue string `url:"-"` //请求头使用，不计入请求内容
}

//发起请求前填充默认信息
func (req *QiniuBaseRequest) B() error {
	return nil
}

//获取请求链接
func (req *QiniuBaseRequest) URL() string {
	return ""
}

//验证请求是否合格
func (req *QiniuBaseRequest) V() (bool, error) {
	return true, nil
}

//设置请求头信息
func (req *QiniuBaseRequest) H() map[string]string {
	headers := map[string]string{
		"Host":         req.host,
		"Content-Type": "application/x-www-form-urlencoded",
	}
	if len(req.authorizationName) > 0 && len(req.AuthorizationValue) > 0 {
		headers["Authorization"] = req.authorizationName + " " + req.AuthorizationValue
	}
	return headers
}

//七牛子账号请求公共参数数据结构
type QiniuBaseChildRequest struct {
	QiniuBaseRequest
	AccessKey string `url:"-"` //请求头使用，不计入请求内容
	SecretKey string `url:"-"` //请求头使用，不计入请求内容
}

//Accesstoken数据结构
type Access struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"-"`
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
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
	req := AccountOauthTokenRequest{}
	req.Username = access.Username
	req.Password = access.Password
	resp := &AccountOauthTokenResponse{}
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

//获取七牛缓存信息保存目录
func GetDir() string {
	// accesspath, _ := gofile.GetUserHome()
	//由于php调用无权限执行根目录 这边根据实际情况先固定以下路径
	accesspath := "/var/www/html"
	if len(accesspath) > 0 {
		return accesspath + string(os.PathSeparator) + ".satool" + string(os.PathSeparator) + "qiniu"
	} else {
		return ".satool" + string(os.PathSeparator) + "qiniu"
	}
}

//使用AK和SK对请求进行签名
func Sign(req gohttp.Request, accesskey, secretkey string) (string, error) {
	params, err := query.Values(req)
	if err != nil {
		return "", err
	} else {
		h := hmac.New(sha1.New, []byte(secretkey))
		u, err := url.Parse(req.URL())
		if err != nil {
			return "", err
		} else {
			data := u.Path
			if u.RawQuery != "" {
				data += "?" + u.RawQuery
				if len(params) > 0 {
					data += "&" + params.Encode()
				}
				data += "\n"
			} else if len(params) > 0 {
				data += "?" + params.Encode() + "\n"
			} else {
				data += "\n"
			}

			h.Write([]byte(data))

			return accesskey + ":" + base64.URLEncoding.EncodeToString([]byte(h.Sum(nil))), nil
		}
	}
}
