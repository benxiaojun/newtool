package newqiniu

// https://itisatest.qnssl.com/kodo/storage.pdf

// Put用量
type UsagePutRequest struct {
	QiniuBaseAKSKRequest
	Begin       string `url:"begin,omitempty"` //"20160601/00:00"
	End         string `url:"end,omitempty"`   //"20160601/00:00"
	Granularity string `url:"g,omitempty"`
	Select      string `url:"select,omitempty"`
	Bucket      string `url:"bucket,omitempty"`
}

func (req *UsagePutRequest) URL() string {
	return "http://" + HOST_API + "/v6/rs_put"
}

func (req *UsagePutRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	req.Select = "hits"
	return nil
}

func (req *UsagePutRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}

// ----------- 存储用量
type UsageSpaceResponse struct {
	QiniuBaseResponse
	Times []int64 `json:"times"`
	Datas []int64 `json:"datas"`
}

type UsageSpaceRequest struct {
	QiniuBaseAKSKRequest
	Begin       string `url:"begin,omitempty"` //"20160731000000"
	End         string `url:"end,omitempty"`   //"20160731000000"
	Granularity string `url:"g,omitempty"`
	Bucket      string `url:"bucket,omitempty"`
}

func (req *UsageSpaceRequest) URL() string {
	return "http://" + HOST_API + "/v6/space"
}

func (req *UsageSpaceRequest) B() error {
	req.host = HOST_API
	req.authorizationName = AUTHORIZATION_QBOX
	return nil
}

func (req *UsageSpaceRequest) V() (bool, error) {
	sign, err := Sign(req, req.AccessKey, req.SecretKey)
	if err != nil {
		return false, err
	} else {
		req.AuthorizationValue = sign
		return true, nil
	}
}
