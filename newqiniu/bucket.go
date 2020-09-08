package newqiniu

import "strconv"

//--------------通过AK/SK创建BUCKET

//通过AK/SK创建BUCKET返回数据结构,请求成功不返回任何内容
type CreateBucketResponse struct {
	QiniuBaseResponse
}

//通过AK/SK创建BUCKET请求数据结构
type CreateBucketRequest struct {
	QiniuBaseAKSKRequest
	Name     string `url:"-"`
	IsPublic int    `url:"-"`
}

//通过AK/SK创建BUCKET请求链接
func (req *CreateBucketRequest) URL() string {
	return "http://" + HOST_RS + "/mkbucket2/" + req.Name + "/public/" + strconv.Itoa(req.IsPublic)
}

//填充默认信息
func (req *CreateBucketRequest) B() error {
	req.host = HOST_RS
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *CreateBucketRequest) V() (bool, error) {
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
type DeleteBucketResponse struct {
	QiniuBaseResponse
}

//通过AK/SK删除BUCKET请求数据结构
type DeleteBucketRequest struct {
	QiniuBaseAKSKRequest
	Name string `url:"-"`
}

//通过AK/SK删除BUCKET请求链接
func (req *DeleteBucketRequest) URL() string {
	return "http://" + HOST_RS + "/drop/" + req.Name
}

//填出默认信息
func (req *DeleteBucketRequest) B() error {
	req.host = HOST_RS
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *DeleteBucketRequest) V() (bool, error) {
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
type ListBucketResponse struct {
	QiniuBaseResponse
}

//通过AK/SK列举BUCKET请求数据结构
type ListBucketRequest struct {
	QiniuBaseAKSKRequest
}

//通过AK/SK列举BUCKET请求链接
func (req *ListBucketRequest) URL() string {
	return "http://" + HOST_RSME + "/buckets"
}

//填充默认字段
func (req *ListBucketRequest) B() error {
	req.host = HOST_RSME
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *ListBucketRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

//--------------通过AK/SK获取BUCKET域名

//通过AK/SK获取BUCKET域名请求数据结构
type ListBucketDomainRequest struct {
	QiniuBaseAKSKRequest
	BucketName string `url:"tbl"`
}

//通过AK/SK获取BUCKET域名请求链接
func (req *ListBucketDomainRequest) URL() string {
	return "http://" + HOST_API + "/v6/domain/list"
}

//填充数据
func (req *ListBucketDomainRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *ListBucketDomainRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

//------------绑定域名

//绑定域名请求数据结构
type BucketBindDomainRequest struct {
	QiniuBaseAKSKRequest
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
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *BucketBindDomainRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

// -------------bucket用量

type BucketStatInfoResponse struct {
	QiniuBaseResponse
	ApicallGet int64 `json:"apicall_get"`
	ApicallPut int64 `json:"apicall_put"`
	Bandwidth  int64 `json:"bandwidth"`
	OvTransfer int64 `json:"ov_transfer"`
	Space      int64 `json:"space"`
	SpaceAvg   int64 `json:"space_avg"`
	Transfer   int64 `json:"transfer"`
}

type BucketStatInfoRequest struct {
	QiniuBaseAKSKRequest
	Uid    int64  `url:"uid,omitempty"`
	Bucket string `url:"bucket,omitempty"`
	Month  string `url:"month,omitempty"`
}

func (req *BucketStatInfoRequest) URL() string {
	return "http://" + HOST_API + "/stat/info"
}

func (req *BucketStatInfoRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

func (req *BucketStatInfoRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}
