package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/sillydong/goczd/godata"
	"net/url"
	"time"
)

const (
	HOST_ECS = "https://ecs.aliyuncs.com"
	HOST_RDS = "https://rds.aliyuncs.com"
)

const TIME_FORMAT = "2006-01-02T15:04:05Z"

type AliyunBaseResponse struct {
	Code      string `json:"Code,omitempty"`
	HostID    string `json:"HostId,omitempty"`
	Message   string `json:"Message,omitempty"`
	RequestID string `json:"RequestId"`
}

func (resp *AliyunBaseResponse) OK() bool {
	return len(resp.Code) == 0
}

func (resp *AliyunBaseResponse) Msg() string {
	return fmt.Sprintf("[%s][%s]%s", resp.RequestID, resp.Code, resp.Message)
}

type AliyunBaseRequest struct {
	Child            interface{} `url:"-"`
	Action           string      `url:"Action"`
	Format           string      `url:"Format"`
	Version          string      `url:"Version"`
	AccessKeyId      string      `url:"AccessKeyId"`
	AccessSecret     string      `url:"-"`
	Signature        string      `url:"Signature,omitempty"`
	SignatureMethod  string      `url:"SignatureMethod"`
	Timestamp        string      `url:"Timestamp"`
	SignatureVersion string      `url:"SignatureVersion"`
	SignatureNonce   string      `url:"SignatureNonce"`
}

func (req *AliyunBaseRequest) B() error {
	req.Format = "JSON"
	req.SignatureMethod = "HMAC-SHA1"
	req.SignatureVersion = "1.0"
	req.SignatureNonce = godata.RandomString(8, godata.ALPHANUMERIC)
	req.Timestamp = time.Now().UTC().Format(TIME_FORMAT)
	return nil
}

func (req *AliyunBaseRequest) V() (bool, error) {
	return true, nil
}

func (req *AliyunBaseRequest) H() map[string]string {
	return nil
}

func (req *AliyunBaseRequest) Sign() error {
	params, err := query.Values(req.Child)
	if err != nil {
		return err
	} else {
		str := "GET&" + percentEncode("/") + "&" + percentEncode(params.Encode())
		h := hmac.New(sha1.New, []byte(req.AccessSecret+"&"))
		h.Write([]byte(str))
		req.Signature = base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
		return nil
	}
}

func percentEncode(content string) string {
	return godata.StrReplace(url.QueryEscape(content), []string{"+", "*", "%7E"}, []string{"%20", "%2A", "~"}, -1)
}
