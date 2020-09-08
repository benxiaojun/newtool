package ecs

import (
	"github.com/benxiaojun/satool/aliyun"
	"time"
)

//获取实例监控数据返回数据结构
type DescribeInstanceMonitorDataResponse struct {
	EcsBaseResponse
	MonitorData struct {
		InstanceMonitorData []InstanceMonitorDataType `json:"InstanceMonitorData"`
	} `json:"MonitorData"`
}

const (
	PERIOD_ONE_MINUTE   = 60   //精度60秒
	PERIOD_TEN_MINUTE   = 600  //精度600秒
	PERIOD_SIXTY_MINUTE = 3600 //精度3600秒
)

//获取实例监控数据请求数据结构
type DescribeInstanceMonitorDataRequest struct {
	EcsBaseRequest
	InstanceId string `url:"InstanceId"`
	StartTime  string `url:"StartTime"`
	EndTime    string `url:"EndTime"`
	Period     int    `url:"Period,omitempty"`
}

//填充默认数据，签名
func (req *DescribeInstanceMonitorDataRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeInstanceMonitorData"
	req.Child = req
	return req.Sign()
}

//设置开始时间，格式2006-01-02 15:04:05
func (req *DescribeInstanceMonitorDataRequest) SetStartTime(data string) {
	time, err := time.Parse("2006-01-02 15:04:05", data)
	if err == nil {
		req.StartTime = time.Format(aliyun.TIME_FORMAT)
	}

}

//设置结束时间，格式2006-01-02 15:04:05
func (req *DescribeInstanceMonitorDataRequest) SetEndTime(data string) {
	time, err := time.Parse("2006-01-02 15:04:05", data)
	if err == nil {
		req.EndTime = time.Format(aliyun.TIME_FORMAT)
	}

}
