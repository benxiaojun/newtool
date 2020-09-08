package ecs

//-----------------查询实例列表

//查询实例列表返回内容数据结构
type DescribeInstancesResponse struct {
	EcsBaseResponse
	TotalCount int64 `json:"TotalCount"`
	PageNumber int64 `json:"PageNumber"`
	PageSize   int64 `json:"PageSize"`
	Instances  struct {
		Instance []InstanceAttributesType
	} `json:"Instances"`
}

//查询实例列表请求数据结构
type DescribeInstancesRequest struct {
	EcsBaseRequest
	RegionId            string `url:"RegionId"`
	VpcId               string `url:"VpcId,omitempty"`
	VSwitchId           string `url:"VSwitchId,omitempty"`
	ZoneId              string `url:"ZoneId,omitempty"`
	InstanceIds         string `url:"InstanceIds,omitempty"`
	InstanceNetworkType string `url:"InstanceNetworkType,omitempty"`
	PrivateIpAddresses  string `url:"PrivateIpAddresses,omitempty"`
	InnerIpAddresses    string `url:"InnerIpAddresses,omitempty"`
	PublicIpAddresses   string `url:"PublicIpAddresses,omitempty"`
	SecurityGroupIds    string `url:"SecurityGroupIds,omitempty"`
	PageNumber          string `url:"PageNumber,omitempty"`
	PageSize            string `url:"PageSize,omitempty"`
}

//填充数据，签名
func (req *DescribeInstancesRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeInstances"
	req.Child = req
	return req.Sign()
}

//---------------查询实例状态

//查询实例状态返回数据结构
type DescribeInstanceStatusResponse struct {
	EcsBaseResponse
	TotalCount       int64                 `json:"TotalCount"`
	PageNumber       int64                 `json:"PageNumber"`
	PageSize         int64                 `json:"PageSize"`
	InstanceStatuses InstanceStatusSetType `json:"InstanceStatuses"`
}

//查询实例状态请求数据结构
type DescribeInstanceStatusRequest struct {
	EcsBaseRequest
	RegionId   string `url:"RegionId"`
	ZoneId     string `url:"ZoneId"`
	PageNumber int    `url:"PageNumber,omitempty"`
	PageSize   int    `url:"PageSize,omitempty"`
}

//填充数据，签名
func (req *DescribeInstanceStatusRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeInstanceStatus"
	req.Child = req
	return req.Sign()
}

//--------------创建实例

//创建实例返回数据结构
type CreateInstanceResponse struct {
	EcsBaseResponse
	InstanceId string `json:"InstanceId"`
}

//创建实例请求数据结构
type CreateInstanceRequest struct {
	EcsBaseRequest
	RegionId                string `url:"RegionId"`
	ZoneId                  string `url:"ZoneId,omitempty"`
	ImageId                 string `url:"ImageId"`
	InstanceType            string `url:"InstanceType"`
	SecurityGroupId         string `url:"SecurityGroupId"`
	InstanceName            string `url:"InsntanceName,omitempty"`
	Description             string `url:"Description,omitempty"`
	InternetChargeType      string `url:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn  string `url:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut string `url:"InternetMaxBandwidthOut,omitempty"`
	HostName                string `url:"HostName,omitempty"`
	Password                string `url:"Passowrd,omitempty"`
	IoOptimized             string `url:"IoOptimized,omitempty"`
	SystemDisk              struct {
		Category    string `url:"SystemDisk.Category,omitempty"`
		DiskName    string `url:"SystemDisk.DiskName,omitempty"`
		Description string `url:"SystemDisk.Description,omitempty"`
	} `url:"-"`
	DataDisk    []CreateInstaceRequestDataDisk `url:"DataDisk,omitempty"`
	VSwitchId   string                         `url:"VSwitchId,omitempty"`
	ClientToken string                         `url:"ClientToken,omitempty"`
}

type CreateInstaceRequestDataDisk struct {
	Size               int64  `url:"Size,omitempty"`
	Category           string `url:"Category,omitempty"`
	SnapshotId         string `url:"SnapshotId,omitempty"`
	DiskName           string `url:"DiskName,omitempty"`
	Description        string `url:"Description,omitempty"`
	Device             string `url:"Device,omitempty"`
	DeleteWithInstance string `url:"DeleteWithInstance,omitempty"`
}

//填充数据，签名
func (req *CreateInstanceRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "CreateInstance"
	req.Child = req
	return req.Sign()
}

type ModifyInstanceAttributeResponse struct {
	EcsBaseResponse
}

type ModifyInstanceAttributeRequest struct {
	EcsBaseRequest
	InstanceId   string `url:"InstanceId"`
	InstanceName string `url:"InstanceName,omitempty"`
	Description  string `url:"Description,omitempty"`
	Password     string `url:"Password,omitempty"`
	HostName     string `url:"HostName,omitempty"`
}

func (req *ModifyInstanceAttributeRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "ModifyInstanceAttribute"
	req.Child = req
	return req.Sign()
}
