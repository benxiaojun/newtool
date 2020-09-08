package dnspod

import "fmt"

//DNSPOD接口请求域名
const HOST = "https://dnsapi.cn"

//DNSPOD返回内容公共数据结构
type DnspodBaseResponse struct {
	Status struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		CreatedAt string `json:"created_at"`
	} `json:"status"`
}

//判断返回内容是否有错
func (resp *DnspodBaseResponse) OK() bool {
	return resp.Status.Code == "1"
}

//获取返回内容信息
func (resp *DnspodBaseResponse) Msg() string {
	return fmt.Sprintf("请求出错:[%s]%s", resp.Status.Code, resp.Status.Message)
}

//DNSPOD请求公共参数数据结构
type DnspodBaseRequest struct {
	LoginToken   string `url:"login_token"`
	Format       string `url:"format,omitempty"`
	Lang         string `url:"lang,omitempty"`
	ErrorOnEmpty string `url:"error_on_empty,omitempty"`
}

//在请求发起前填充默认数据
func (req *DnspodBaseRequest) B() error {
	req.Format = "json"
	req.Lang = "cn"
	req.ErrorOnEmpty = "yes"
	return nil
}

//默认请求地址
func (req *DnspodBaseRequest) URL() string {
	return ""
}

//默认请求校验函数，返回true
func (req *DnspodBaseRequest) V() (bool, error) {
	return true, nil
}

//默认设置请求头，返回nil
func (req *DnspodBaseRequest) H() map[string]string {
	return nil
}
