package qiniu

import (
	"encoding/json"
	"fmt"
	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gohttp"
	"os"
	"strconv"
	"time"
)

//------------------创建子账号

//创建子账号返回数据结构
type CreateChildAccountResponse struct {
	QiniuBaseResponse
	DeviceNum             int    `json:"device_num,omitempty"`
	Email                 string `json:"email"`
	InvitationNum         int    `json:"invitation_num,omitempty"`
	IsActivatied          bool   `json:"is_activated,omitempty"`
	IsDisabled            bool   `json:"is_disabled,omitempty"`
	LastParentOperationAt string `json:"last_parent_operation_at"`
	ParentUID             int    `json:"parent_uid"`
	UID                   int64  `json:"uid"`
	UserType              int    `json:"user_type,omitempty"`
	Userid                string `json:"userid"`
}

//创建子账号请求数据结构
type CreateChildAccountRequest struct {
	QiniuBaseRequest
	Email    string `url:"email"`
	Password string `url:"password"`
}

//创建子账号请求链接
func (req *CreateChildAccountRequest) URL() string {
	return "https://" + HOST_ACC + "/user/create_child"
}

//填充默认数据
func (req *CreateChildAccountRequest) B() error {
	req.host = HOST_ACC
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//------------------列举子账号

//列举子账号返回数据结构
type ListChildAccountResponse struct {
	QiniuBaseResponse
	DeviceNum             int64     `json:"device_num"`
	Email                 string    `json:"email"`
	InvitationNum         int64     `json:"invitation_num"`
	IsActivated           bool      `json:"is_activated"`
	IsDisabled            bool      `json:"is_disabled"`
	LastParentOperationAt time.Time `json:"last_parent_operation_at"`
	ParentUid             int64     `json:"parent_uid"`
	Uid                   int64     `json:"uid"`
	UserType              int64     `json:"user_type"`
	Userid                string    `json:"userid"`
}

//列举子账号请求数据结构
type ListChildAccountRequest struct {
	QiniuBaseRequest
	Offset int `url:"offset"`
	Limit  int `url:"limit"`
}

//列举子账号请求链接
func (req *ListChildAccountRequest) URL() string {
	return "https://" + HOST_ACC + "/user/children"
}

//填充默认数据
func (req *ListChildAccountRequest) B() error {
	req.host = HOST_ACC
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//------------------禁用子账号
//禁用子账号返回数据结构
type DisableChildAccountResponse struct {
	QiniuBaseResponse
	DeviceNum             int64     `json:"device_num"`
	Email                 string    `json:"email"`
	InvitationNum         int64     `json:"invitation_num"`
	IsActivated           bool      `json:"is_activated"`
	IsDisabled            bool      `json:"is_disabled"`
	LastParentOperationAt time.Time `json:"last_parent_operation_at"`
	ParentUid             int64     `json:"parent_uid"`
	Uid                   int64     `json:"uid"`
	UserType              int64     `json:"user_type"`
	Userid                string    `json:"userid"`
}

//禁用子账号请求数据结构
type DisableChildAccountRequest struct {
	QiniuBaseRequest
	Uid    int64  `url:"uid"`
	Reason string `url:"reason"`
}

//禁用子账号请求链接
func (req *DisableChildAccountRequest) URL() string {
	return "https://" + HOST_ACC + "/user/disable_child"
}

//填充默认数据
func (req *DisableChildAccountRequest) B() error {
	req.host = HOST_ACC
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//------------------启用子账号
//启用子账号返回数据结构
type EnableChildAccountResponse struct {
	QiniuBaseResponse
	DeviceNum             int64     `json:"device_num"`
	Email                 string    `json:"email"`
	InvitationNum         int64     `json:"invitation_num"`
	IsActivated           bool      `json:"is_activated"`
	IsDisabled            bool      `json:"is_disabled"`
	LastParentOperationAt time.Time `json:"last_parent_operation_at"`
	ParentUid             int64     `json:"parent_uid"`
	Uid                   int64     `json:"uid"`
	UserType              int64     `json:"user_type"`
	Userid                string    `json:"userid"`
}

//启用子账号请求数据结构
type EnableChildAccountRequest struct {
	QiniuBaseRequest
	Uid int64 `url:"uid"`
}

//启用子账号请求链接
func (req *EnableChildAccountRequest) URL() string {
	return "https://" + HOST_ACC + "/user/enable_child"
}

//填充默认数据
func (req *EnableChildAccountRequest) B() error {
	req.host = HOST_ACC
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//------------------获取账号AK/SK

//获取账户AK/SK返回数据结构
type ChildKeysResponse struct {
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

//获取子账户AK/SK请求数据结构
type ChildKeysRequest struct {
	QiniuBaseRequest
	Uid   int64  `url:"uid,omitempty"`
	Email string `url:"email,omitempty"`
}

//获取子账户AK/SK请求链接
func (req *ChildKeysRequest) URL() string {
	return "https://" + HOST_ACC + "/user/child_key"
}

//填充默认数据
func (req *ChildKeysRequest) B() error {
	req.host = HOST_ACC
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//判断请求参数
func (req *ChildKeysRequest) V() (bool, error) {
	if req.Uid > 0 || req.Email != "" {
		return true, nil
	}
	return false, fmt.Errorf("Uid或Emial需提供其中一个")
}

type Child struct {
	Uid       int64  `json:"uid"`
	Key       string `json:"key"`
	Key2      string `json:"key2"`
	Secret    string `json:"secret"`
	Secret2   string `json:"secret2"`
	ExpiresIn int64  `json:"expires_in"`
}

func (child *Child) GetFileName() string {
	return strconv.FormatInt(child.Uid, 10) + ".json"
}

func (child *Child) IsExpired() bool {
	return child.ExpiresIn <= (time.Now().Unix() - 60)
}

func (child *Child) Get(accesstoken string) error {
	if child.Uid > 0 {
		req := ChildKeysRequest{}
		req.AuthorizationValue = accesstoken
		req.Uid = child.Uid
		resp := &ChildKeysResponse{}
		err := gohttp.DoGetResponse(&req, resp)
		if err != nil {
			return err
		} else {
			if resp.OK() {
				child.Key = resp.Key
				child.Key2 = resp.Key2
				child.Secret = resp.Secret
				child.Secret2 = resp.Secret2
				child.ExpiresIn = time.Now().Unix() + 86400
				return child.WriteKeys()
			} else {
				return fmt.Errorf("%v", resp.Msg())
			}
		}
	} else {
		return fmt.Errorf("未设置子账号ID")
	}
}

func (child *Child) Refresh(accesstoken string) error {
	return child.Get(accesstoken)
}

func (child *Child) WriteKeys() error {
	if bytes, err := json.Marshal(child); err != nil {
		return err
	} else {
		if accesspath := GetDir(); len(accesspath) == 0 {
			return err
		} else {
			err := os.MkdirAll(accesspath, 0777)
			if err != nil {
				return err
			}
			if err := gofile.Json2File(string(bytes), accesspath+string(os.PathSeparator)+child.GetFileName(), 0777); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}

func (child *Child) ReadKeys(autorefresh bool, accesstoken string) error {
	childpath := GetDir() + string(os.PathSeparator) + child.GetFileName()
	if err := gofile.File2Json(childpath, &child); err != nil {
		return err
	} else {
		if child.IsExpired() && autorefresh {
			return child.Refresh(accesstoken)
		}
		return nil
	}
}
