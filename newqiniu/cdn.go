package newqiniu

// http://developer.qiniu.com/article/fusion/api/traffic-bandwidth.html

const (
	GRANULARITY_5MIN = "5min"
	GRANULARITY_HOUR = "hour"
	GRANULARITY_DAY  = "day"
)

//------------域名带宽
//通过AK/SK获取域名带宽结构
type DomainBandwidthResponse struct {
	QiniuBaseResponse
	Time []string                      `json:"time"`
	Data map[string]map[string][]int64 `json:"data"`
}

//通过AK/SK获取BUCKET域名带宽请求数据结构
type DomainBandwidthRequest struct {
	QiniuBaseAKSKRequest
	Domains     string `url:"-" json:"domains"`
	StartDate   string `url:"-" json:"startDate"` //"2016-07-01"
	EndDate     string `url:"-" json:"endDate"`   //"2016-07-01"
	Granularity string `url:"-" json:"granularity"`
}

//通过AK/SK获取域名带宽请求链接
func (req *DomainBandwidthRequest) URL() string {
	return "http://" + HOST_FUSION + "/v2/tune/bandwidth"
}

//填充数据
func (req *DomainBandwidthRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *DomainBandwidthRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

func (req *DomainBandwidthRequest) H() map[string]string {
	headers := req.QiniuBaseAKSKRequest.H()
	headers["Content-Type"] = "application/json"
	return headers
}

//------------域名流量
//通过AK/SK获取域名流量结构
type DomainTrafficResponse struct {
	QiniuBaseResponse
	Time []string                      `json:"time"`
	Data map[string]map[string][]int64 `json:"data"`
}

//通过AK/SK获取BUCKET域名请求数据结构
type DomainTrafficRequest struct {
	QiniuBaseAKSKRequest
	Domains     string `url:"-" json:"domains"`
	StartDate   string `url:"-" json:"startDate"`
	EndDate     string `url:"-" json:"endDate"`
	Granularity string `url:"-" json:"granularity"`
}

//通过AK/SK获取域名流量请求链接
func (req *DomainTrafficRequest) URL() string {
	return "http://" + HOST_FUSION + "/v2/tune/flux"
}

//填充数据
func (req *DomainTrafficRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

//请求签名
func (req *DomainTrafficRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

func (req *DomainTrafficRequest) H() map[string]string {
	headers := req.QiniuBaseAKSKRequest.H()
	headers["Content-Type"] = "application/json"
	return headers
}
