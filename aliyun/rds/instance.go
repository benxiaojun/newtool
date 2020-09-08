package rds

type DescribeDBInstancesResponse struct {
	RdsBaseResponse
	TotalRecordCount int64 `json:"TotalRecordCount"`
	PageNumber       int64 `json:"PageNumber"`
	PageRecordCount  int64 `json:"PageRecordCount"`
	Items            struct {
		DBInstance []DBInstance `json:"DBInstance"`
	} `json:"Items"`
}

type DescribeDBInstancesRequest struct {
	RdsBaseRequest
	RegionId            string `url:"RegionId"`
	Engine              string `url:"Engine,omitempty"`
	DBInstanceType      string `url:"DbInstanceType,omitempty"`
	InstanceNetworkType string `url:"InstanceNetworkType,omitempty"`
	ConnectionMode      string `url:"ConnectionMode,omitempty"`
	PageSize            int    `url:"PageSize,omitempty"`
	PageNumber          int    `url:"PageNumber,omitempty"`
}

func (req *DescribeDBInstancesRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeDBInstances"
	req.Child = req
	return req.Sign()
}

type ModifyDBInstanceDescriptionResponse struct {
	RdsBaseResponse
}

type ModifyDBInstanceDescriptionRequest struct {
	RdsBaseRequest
	DBInstanceId          string `url:"DBInstanceId"`
	DBInstanceDescription string `url:"DBInstanceDescription"`
}

func (req *ModifyDBInstanceDescriptionRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "ModifyDBInstanceDescription"
	req.Child = req
	return req.Sign()
}

type DescribeDBInstanceAttributeResponse struct {
	RdsBaseResponse
	Items struct {
		DBInstanceAttribute []DBInstanceAttribute `json:"DBInstanceAttribute"`
	} `json:"Items"`
}

type DescribeDBInstanceAttributeRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
}

func (req *DescribeDBInstanceAttributeRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeDBInstanceAttribute"
	req.Child = req
	return req.Sign()
}

type RestartDBInstanceResponse struct {
	RdsBaseResponse
}

type RestartDBInstanceRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
}

func (req RestartDBInstanceRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "RestartDBInstance"
	req.Child = req
	return req.Sign()
}

type ModifySecurityIpsResponse struct {
	RdsBaseResponse
}

type ModifySecurityIpsRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	SecurityIps  string `url:"SecurityIps"`
	DBInstanceIPArrayName  string `url:"DBInstanceIPArrayName"`
	DBInstanceIPArrayAttribute  string `url:"DBInstanceIPArrayAttribute"`
}

func (req *ModifySecurityIpsRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "ModifySecurityIps"
	req.Child = req
	return req.Sign()
}

type UpgradeDBInstanceEngineVersionResponse struct {
	RdsBaseResponse
	TaskId int64 `json:"TaskId,string"`
}

type UpgradeDBInstanceEngineVersionRequest struct {
	RdsBaseRequest
	DBInstanceId  string `url:"DBInstanceId"`
	EngineVersion string `url:"EngineVersion"`
}

func (req *UpgradeDBInstanceEngineVersionRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "UpgradeDBInstanceEngineVersion"
	req.Child = req
	return req.Sign()
}
