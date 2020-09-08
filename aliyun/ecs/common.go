package ecs

import (
	"github.com/benxiaojun/satool/aliyun"
)

//阿里云ECS返回内容公共数据结构
type EcsBaseResponse struct {
	aliyun.AliyunBaseResponse
}

type EcsBaseRequest struct {
	aliyun.AliyunBaseRequest
}

func (req *EcsBaseRequest) URL() string {
	return aliyun.HOST_ECS
}

func (req *EcsBaseRequest) B() error {
	req.Version = "2014-05-26"
	return req.AliyunBaseRequest.B()
}
