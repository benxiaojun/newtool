package qiniu

import "strconv"

//--------------创建BUCKET

//创建BUCKET返回数据结构,请求成功不返回任何内容
type CreateBucketResponse struct {
	QiniuBaseResponse
}

//创建BUCKET请求数据结构
type CreateBucketRequest struct {
	QiniuBaseRequest
	Name     string `url:"-"`
	IsPublic int    `url:"-"`
}

//创建BUCKET请求链接
func (req *CreateBucketRequest) URL() string {
	return "http://" + HOST_RS + "/mkbucket2/" + req.Name + "/public/" + strconv.Itoa(req.IsPublic)
}

//填充默认信息
func (req *CreateBucketRequest) B() error {
	req.host = HOST_RS
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//--------------删除BUCKET

//删除BUCKET返回数据结构,请求成功不返回任何内容
type DeleteBucketResponse struct {
	QiniuBaseResponse
}

//删除BUCKET请求数据结构
type DeleteBucketRequest struct {
	QiniuBaseRequest
	Name string `url:"-"`
}

//删除BUCKET请求链接
func (req *DeleteBucketRequest) URL() string {
	return "http://" + HOST_RS + "/drop/" + req.Name
}

//填出默认信息
func (req *DeleteBucketRequest) B() error {
	req.host = HOST_RS
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//--------------列举BUCKET

//列举BUCKET返回数据结构
type ListBucketResponse struct {
	QiniuBaseResponse
}

//列举BUCKET请求数据结构
type ListBucketRequest struct {
	QiniuBaseRequest
}

//列举BUCKET请求链接
func (req *ListBucketRequest) URL() string {
	return "http://" + HOST_RSME + "/buckets"
}

//填充默认字段
func (req *ListBucketRequest) B() error {
	req.host = HOST_RSME
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//--------------获取BUCKET域名

//获取BUCKET域名返回数据结构
type ListBucketDomainResponse struct {
	QiniuBaseResponse
}

//获取BUCKET域名请求数据结构
type ListBucketDomainRequest struct {
	QiniuBaseRequest
	BucketName string `url:"tbl"`
}

//获取BUCKET域名请求链接
func (req *ListBucketDomainRequest) URL() string {
	return "http://" + HOST_API + "/v6/domain/list"
}

//填充数据
func (req *ListBucketDomainRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}

//--------------绑定域名

//绑定域名返回数据结构
type BucketBindDomainResponse struct {
	QiniuBaseResponse
}

//绑定域名请求数据结构
type BucketBindDomainRequest struct {
	QiniuBaseRequest
	Domain     string `url:"domain"`
	BucketName string `url:"tbl"`
}

//绑定域名请求链接
func (req *BucketBindDomainRequest) URL() string {
	return "http://" + HOST_API + "/v6/domain/publish"
}

//填充数据
func (req *BucketBindDomainRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_BEARER
	return nil
}
