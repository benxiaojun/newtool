package rds

import "github.com/benxiaojun/satool/aliyun"

type RdsBaseResponse struct {
	aliyun.AliyunBaseResponse
}

type RdsBaseRequest struct {
	aliyun.AliyunBaseRequest
}

func (req *RdsBaseRequest) URL() string {
	return aliyun.HOST_RDS
}

func (req *RdsBaseRequest) B() error {
	req.Version = "2014-08-15"
	return req.AliyunBaseRequest.B()
}
