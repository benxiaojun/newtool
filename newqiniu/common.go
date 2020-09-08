package newqiniu

import (
	"os"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gohttp"
)

const (
	HOST_FUSION = "fusion.qiniuapi.com"
	HOST_ACC    = "acc.qbox.me"
	HOST_UC     = "uc.qbox.me"
	HOST_RS     = "rs.qiniu.com"
	HOST_RSME   = "rs.qbox.me"
	HOST_API    = "api.qiniu.com"
)

const (
	BUCKET_PRIVATE = 0
	BUCKET_PUBLIC  = 1
)

//请求公共返回内容数据结构
type QiniuBaseResponse struct {
	Code             int    `json:"code,omitempty"`
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
	}

	return fmt.Sprintf("[%d]%s:%s", resp.ErrorCode, resp.Error, resp.ErrorDescription)
}

const (
	AUTHORIZATION_BEARER = "Bearer" //主账号使用Authorization字段
	AUTHORIZATION_QBOX   = "QBox"   //子账号使用Authorization字段
)

//七牛请求公共参数数据结构
type QiniuBaseRequest struct {
	host               string `url:"-" json:"-"` //请求头使用，不计入请求内容
	authorizationName  string `url:"-" json:"-"` //请求头使用，不计入请求内容
	AuthorizationValue string `url:"-" json:"-"` //请求头使用，不计入请求内容
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
type QiniuBaseAKSKRequest struct {
	QiniuBaseRequest
	AccessKey string `url:"-" json:"-"` //请求头使用，不计入请求内容
	SecretKey string `url:"-" json:"-"` //请求头使用，不计入请求内容
}

//获取七牛缓存信息保存目录
func GetDir() string {
	accesspath, _ := gofile.GetUserHome()
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
