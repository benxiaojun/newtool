package qiniu

import (
	// "strconv"
)

//--------------通过AK/SK创建BUCKET

//通过AK/SK创建BUCKET返回数据结构,请求成功不返回任何内容
type KeyCreateBucketResponse struct {
	QiniuBaseResponse
}

//通过AK/SK创建BUCKET请求数据结构
type KeyCreateBucketRequest struct {
	QiniuBaseChildRequest
	Name     string `url:"-"`
	IsPublic int    `url:"-"`
}

//通过AK/SK创建BUCKET请求链接
func (req *KeyCreateBucketRequest) URL() string {
	return "http://" + HOST_RS + "/mkbucketv3/" + req.Name + "/region/z0" 
}

//填充默认信息
func (req *KeyCreateBucketRequest) B() error {
	req.host = HOST_RS
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *KeyCreateBucketRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

//--------------通过AK/SK删除BUCKET

//通过AK/SK删除BUCKET返回数据结构
type KeyDeleteBucketResponse struct {
	QiniuBaseResponse
}

//通过AK/SK删除BUCKET请求数据结构
type KeyDeleteBucketRequest struct {
	QiniuBaseChildRequest
	Name string `url:"-"`
}

//通过AK/SK删除BUCKET请求链接
func (req *KeyDeleteBucketRequest) URL() string {
	return "http://" + HOST_RS + "/drop/" + req.Name
}

//填出默认信息
func (req *KeyDeleteBucketRequest) B() error {
	req.host = HOST_RS
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *KeyDeleteBucketRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

//--------------通过AK/SK列举BUCKET

//通过AK/SK列举BUCKET返回数据结构
type KeyListBucketResponse struct {
	QiniuBaseResponse
}

//通过AK/SK列举BUCKET请求数据结构
type KeyListBucketRequest struct {
	QiniuBaseChildRequest
}

//通过AK/SK列举BUCKET请求链接
func (req *KeyListBucketRequest) URL() string {
	return "http://" + HOST_RSME + "/buckets"
}

//填充默认字段
func (req *KeyListBucketRequest) B() error {
	req.host = HOST_RSME
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *KeyListBucketRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

//--------------通过AK/SK获取BUCKET域名

//通过AK/SK获取BUCKET域名返回数据结构
type KeyListBucketDomainResponse struct {
	QiniuBaseResponse
}

//通过AK/SK获取BUCKET域名请求数据结构
type KeyListBucketDomainRequest struct {
	QiniuBaseChildRequest
	BucketName string `url:"tbl"`
}

//通过AK/SK获取BUCKET域名请求链接
func (req *KeyListBucketDomainRequest) URL() string {
	return "http://" + HOST_API + "/v6/domain/list"
}

//填充数据
func (req *KeyListBucketDomainRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *KeyListBucketDomainRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

//------------绑定域名

//绑定域名返回数据结构
type KeyBucketBindDomainResponse struct {
	QiniuBaseResponse
}

//绑定域名请求数据结构
type KeyBucketBindDomainRequest struct {
	QiniuBaseChildRequest
	Domain     string `url:"domain"`
	BucketName string `url:"tbl"`
}

//绑定域名请求链接
func (req *KeyBucketBindDomainRequest) URL() string {
	return "http://" + HOST_API + "/v6/domain/publish"
}

//填充数据
func (req *KeyBucketBindDomainRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *KeyBucketBindDomainRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

type ChildStatInfoResponse struct {
	QiniuBaseResponse
	ApicallGet int64 `json:"apicall_get"`
	ApicallPut int64 `json:"apicall_put"`
	Bandwidth  int64 `json:"bandwidth"`
	OvTransfer int64 `json:"ov_transfer"`
	Space      int64 `json:"space"`
	SpaceAvg   int64 `json:"space_avg"`
	Transfer   int64 `json:"transfer"`
}

type ChildStatInfoRequest struct {
	QiniuBaseChildRequest
	Uid    int64  `url:"uid,omitempty"`
	Bucket string `url:"bucket,omitempty"`
	Month  string `url:"month,omitempty"`
}

func (req *ChildStatInfoRequest) URL() string {
	return "http://" + HOST_API + "/stat/info"
}

func (req *ChildStatInfoRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

func (req *ChildStatInfoRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}
