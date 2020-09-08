package dnspod

import (
	"fmt"
	"github.com/sillydong/goczd/godata"
)

//----------------DomainList

//域名列表返回内容数据结构
type DomainListResponse struct {
	DnspodBaseResponse
	Domains []struct {
		AuthToAnquanbao  bool   `json:"auth_to_anquanbao"`
		CnameSpeedup     string `json:"cname_speedup"`
		CreatedOn        string `json:"created_on"`
		ExtStatus        string `json:"ext_status"`
		Grade            string `json:"grade"`
		GradeTitle       string `json:"grade_title"`
		GroupID          string `json:"group_id"`
		ID               int    `json:"id"`
		IsMark           string `json:"is_mark"`
		IsVip            string `json:"is_vip"`
		Name             string `json:"name"`
		Owner            string `json:"owner"`
		Punycode         string `json:"punycode"`
		Records          string `json:"records"`
		Remark           string `json:"remark"`
		SearchenginePush string `json:"searchengine_push"`
		Status           string `json:"status"`
		TTL              string `json:"ttl"`
		UpdatedOn        string `json:"updated_on"`
	} `json:"domains"`
	Info struct {
		AllTotal      int `json:"all_total"`
		DomainTotal   int `json:"domain_total"`
		ErrorTotal    int `json:"error_total"`
		IsmarkTotal   int `json:"ismark_total"`
		LockTotal     int `json:"lock_total"`
		MineTotal     int `json:"mine_total"`
		PauseTotal    int `json:"pause_total"`
		ShareOutTotal int `json:"share_out_total"`
		ShareTotal    int `json:"share_total"`
		SpamTotal     int `json:"spam_total"`
		VipExpire     int `json:"vip_expire"`
		VipTotal      int `json:"vip_total"`
	} `json:"info"`
}

//域名类型数组
var DOMAIN_LIST_TYPES = []string{
	DOMAIN_LIST_TYPE_ALL,
	DOMAIN_LIST_TYPE_MINE,
	DOMAIN_LIST_TYPE_SHARE,
	DOMAIN_LIST_TYPE_ISMARK,
	DOMAIN_LIST_TYPE_PAUSE,
	DOMAIN_LIST_TYPE_VIP,
	DOMAIN_LIST_TYPE_RECENT,
	DOMAIN_LIST_TYPE_SHAREOUT,
}

//域名列表可用域名类型
const (
	DOMAIN_LIST_TYPE_ALL      = "all"       //所有域名
	DOMAIN_LIST_TYPE_MINE     = "mine"      //我的域名
	DOMAIN_LIST_TYPE_SHARE    = "share"     //共享给我的域名
	DOMAIN_LIST_TYPE_ISMARK   = "ismark"    //星标域名
	DOMAIN_LIST_TYPE_PAUSE    = "pause"     //暂停域名
	DOMAIN_LIST_TYPE_VIP      = "vip"       //VIP域名
	DOMAIN_LIST_TYPE_RECENT   = "recent"    //最近操作过的域名
	DOMAIN_LIST_TYPE_SHAREOUT = "share_out" //我共享出去的域名
)

//域名列表请求数据结构
//
//如果账户中的域名数量超过了3000, 将会强制分页并且只返回前3000条, 这时需要通过 offset 和 length 参数去获取其它域名.
type DomainListRequest struct {
	DnspodBaseRequest
	Type    string `url:"type,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Length  int    `url:"length,omitempty"`
	GroupID string `url:"group_id,omitempty"`
	Keyword string `url:"keyword,omitempty"`
}

//生成域名列表请求URL
func (req *DomainListRequest) URL() string {
	return HOST + "/Domain.List"
}

//判断域名列表内容是否正确
func (req *DomainListRequest) V() (bool, error) {
	if len(req.Type) > 0 {
		if godata.InArray(DOMAIN_LIST_TYPES, req.Type) {
			return true, nil
		} else {
			return false, fmt.Errorf("域名类型错误")
		}
	} else {
		return true, nil
	}
}

//----------------RecordType

//获取等级允许的记录类型返回内容数据结构
type RecordTypeResponse struct {
	DnspodBaseResponse
	Types []string `json:"types"`
}

//域名等级数组
var DOMAIN_GRADES = []string{
	DOMAIN_GRADE_D_FREE,
	DOMAIN_GRADE_D_PLUS,
	DOMAIN_GRADE_D_EXTRA,
	DOMAIN_GRADE_D_EXPERT,
	DOMAIN_GRADE_D_ULTRA,
	DOMAIN_GRADE_DP_FREE,
	DOMAIN_GRADE_DP_PLUS,
	DOMAIN_GRADE_DP_EXTRA,
	DOMAIN_GRADE_DP_EXPERT,
	DOMAIN_GRADE_DP_ULTRA,
}

//域名等级
const (
	DOMAIN_GRADE_D_FREE    = "D_Free"    //免费套餐
	DOMAIN_GRADE_D_PLUS    = "D_Plus"    //个人豪华套餐
	DOMAIN_GRADE_D_EXTRA   = "D_Extra"   //企业1
	DOMAIN_GRADE_D_EXPERT  = "D_Expert"  //企业2
	DOMAIN_GRADE_D_ULTRA   = "D_Ultra"   //企业3
	DOMAIN_GRADE_DP_FREE   = "DP_Free"   //新免费套餐
	DOMAIN_GRADE_DP_PLUS   = "DP_Plus"   //新个人专业版
	DOMAIN_GRADE_DP_EXTRA  = "DP_Extra"  //新企业创业版
	DOMAIN_GRADE_DP_EXPERT = "DP_Expert" //新企业标准版
	DOMAIN_GRADE_DP_ULTRA  = "DP_Ultra"  //新企业旗舰版
)

//获取等级允许的记录类型请求数据结构
type RecordTypeRequest struct {
	DnspodBaseRequest
	DomainGrade string `url:"domain_grade"`
}

//获取请求URL
func (req *RecordTypeRequest) URL() string {
	return HOST + "/Record.Type"
}

//判断域名等级是否正确
func (req *RecordTypeRequest) V() (bool, error) {
	if godata.InArray(DOMAIN_GRADES, req.DomainGrade) {
		return true, nil
	} else {
		return false, fmt.Errorf("域名等级错误")
	}
}

//------------------RecordLine

//获取等级允许的线路返回数据结构
type RecordLineResponse struct {
	DnspodBaseResponse
	Lines []string `json:"lines"`
}

//获取等级允许的线路请求数据结构
//
//domain_id 或 domain, 分别对应域名ID和域名, 提交其中一个即可
type RecordLineRequest struct {
	DnspodBaseRequest
	DomainGrade string `url:"domain_grade"`
	DomainId    string `url:"domain_id,omitempty"`
	Domain      string `url:"domain,omitempty"`
}

//获取请求URL
func (req *RecordLineRequest) URL() string {
	return HOST + "/Record.Line"
}

//判断域名等级是否正确，判断是否提供DomainID或Domain
func (req *RecordLineRequest) V() (bool, error) {
	if godata.InArray(DOMAIN_GRADES, req.DomainGrade) {
		if len(req.DomainId) == 0 && len(req.Domain) == 0 {
			return false, fmt.Errorf("DomainId或Domain必须提交其中一个")
		} else {
			return true, nil
		}
	} else {
		return false, fmt.Errorf("域名等级错误")
	}
}
