package dnspod

import "fmt"

//--------------------RecordCreate

//添加记录返回数据结构
type RecordCreateResponse struct {
	DnspodBaseResponse
	Record struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"record"`
}

//添加记录请求数据结构
type RecordCreateRequest struct {
	DnspodBaseRequest
	DomainId   string `url:"domain_id"`
	SubDomain  string `url:"sub_domain"`
	RecordType string `url:"record_type"`
	RecordLine string `url:"record_line"`
	Value      string `url:"value"`
	Mx         int    `url:"mx,omitempty"`
	Ttl        int    `url:"ttl,omitempty"`
	Status     string `url:"status,omitempty"`
}

//获取请求地址
func (req *RecordCreateRequest) URL() string {
	return HOST + "/Record.Create"
}

//判断类型为MX时是否提供Mx优先级
func (req *RecordCreateRequest) V() (bool, error) {
	if req.RecordType == "MX" && (req.Mx < 1 || req.Mx > 20) {
		return false, fmt.Errorf("请提供MX优先级，范围1-20")
	} else {
		return true, nil
	}
}

//-------------RecordList

//记录列表返回数据结构
type RecordListResponse struct {
	DnspodBaseResponse
	Domain struct {
		Grade    string `json:"grade"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Owner    string `json:"owner"`
		Punycode string `json:"punycode"`
	} `json:"domain"`
	Info struct {
		RecordTotal interface{} `json:"record_total"`
		SubDomains  string      `json:"sub_domains"`
	} `json:"info"`
	Records []struct {
		Enabled       string `json:"enabled"`
		ID            string `json:"id"`
		Line          string `json:"line"`
		MonitorStatus string `json:"monitor_status"`
		Mx            string `json:"mx"`
		Name          string `json:"name"`
		Remark        string `json:"remark"`
		Status        string `json:"status"`
		TTL           string `json:"ttl"`
		Type          string `json:"type"`
		UpdatedOn     string `json:"updated_on"`
		UseAqb        string `json:"use_aqb"`
		Value         string `json:"value"`
	} `json:"records"`
}

//记录列表请求数据结构
//
//如果域名的记录数量超过了3000，将会强制分页并且只返回前3000条，这时需要通过 offset 和 length 参数去获取其它记录。
type RecordListRequest struct {
	DnspodBaseRequest
	DomainId  string `url:"domain_id"`
	Offset    int    `url:"offset,omitempty"`
	Length    int    `url:"length,omitempty"`
	SubDomain string `url:"sub_domain,omitempty"`
	Keyword   string `url:"keyword,omitempty"`
}

//请求链接
func (req *RecordListRequest) URL() string {
	return HOST + "/Record.List"
}

//判断请求内容
func (req *RecordListRequest) V() (bool, error) {
	if len(req.DomainId) == 0 {
		return false, fmt.Errorf("DomainId不能为空")
	} else {
		return true, nil
	}
}

//--------------------RecordModify

//修改记录返回数据结构
type RecordModifyResponse struct {
	DnspodBaseResponse
	Record struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
		Value  string `json:"value"`
	} `json:"record"`
}

//修改记录请求数据结构
//
//如果1小时之内，提交了超过5次没有任何变动的记录修改请求，该记录会被系统锁定1小时，不允许再次修改。比如原记录值已经是 1.1.1.1，新的请求还要求修改为 1.1.1.1。
type RecordModifyRequest struct {
	DnspodBaseRequest
	DomainId   string `url:"domain_id"`
	RecordId   string `url:"record_id"`
	SubDomain  string `url:"sub_domain,omitempty"`
	RecordType string `url:"record_type"`
	RecordLine string `url:"record_line"`
	Value      string `url:"value"`
	Mx         int    `url:"mx,omitempty"`
	Ttl        int    `url:"ttl,omitempty"`
	Status     string `url:"status,omitempty"`
}

//请求链接
func (req *RecordModifyRequest) URL() string {
	return HOST + "/Record.Modify"
}

//---------------RecordRemove

//删除记录返回数据结构
type RecordRemoveResponse struct {
	DnspodBaseResponse
}

//删除记录请求数据结构
type RecordRemoveRequest struct {
	DnspodBaseRequest
	DomainId string `url:"domain_id,omitempty"`
	Domain   string `url:"domain,omitempty"`
}

//请求链接
func (req *RecordRemoveRequest) URL() string {
	return HOST + "/Record.Remove"
}

//判断是否提供DomainId或Domain
func (req *RecordRemoveRequest) V() (bool, error) {
	if len(req.DomainId) == 0 && len(req.Domain) == 0 {
		return false, fmt.Errorf("DomainId或Domain必须提交其中一个")
	} else {
		return true, nil
	}
}
