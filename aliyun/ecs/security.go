package ecs

//-----------------查询安全组列表

//查询安全组列表返回数据结构
type DescribeSecurityGroupsResponse struct {
	EcsBaseResponse
	TotalCount     int64  `json:"TotalCount"`
	PageNumber     int64  `json:"PageNumber"`
	PageSize       int64  `json:"PageSize"`
	RegionId       string `json:"RegionId"`
	SecurityGroups struct {
		SecurityGroup []SecurityGroupItemType `json:"SecurityGroup"`
	} `json:"SecurityGroups"`
}

//查询安全组列表请求数据结构
type DescribeSecurityGroupsRequest struct {
	EcsBaseRequest
	RegionId   string `url:"RegionId"`
	VpcId      string `url:"VpcId,omitempty"`
	PageNumber int64  `url:"PageNumber,omitempty"`
	PageSize   int64  `url:"PageSize,omitempty"`
}

//填充数据，签名
func (req *DescribeSecurityGroupsRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeSecurityGroups"
	req.Child = req
	return req.Sign()
}

//-------------查询安全组规则

//查询安全组规则返回数据结构
type DescribeSecurityGroupAttributeResponse struct {
	EcsBaseResponse
	Description       string            `json:"Description"`
	Permissions       PermissionSetType `json:"Permissions"`
	RegionID          string            `json:"RegionId"`
	SecurityGroupID   string            `json:"SecurityGroupId"`
	SecurityGroupName string            `json:"SecurityGroupName"`
	VpcID             string            `json:"VpcId"`
}

//查询安全组规则请求数据结构
type DescribeSecurityGroupAttributeRequest struct {
	EcsBaseRequest
	SecurityGroupId string `url:"SecurityGroupId"`
	RegionId        string `url:"RegionId"`
	NicType         string `url:"NicType,omitempty"` //网络类型
}

//填充数据，签名
func (req *DescribeSecurityGroupAttributeRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeSecurityGroupAttribute"
	req.Child = req
	return req.Sign()
}
