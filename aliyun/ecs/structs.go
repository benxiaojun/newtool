package ecs

import "time"

var DISK_TYPES = map[string]string{
	"system": "系统盘",
	"data":   "数据盘",
}

type DiskItemType struct {
	DiskId           string `json:"DiskId"`
	RegionId         string `json:"RegionId"`
	ZoneId           string `json:"ZoneId"`
	DiskName         string `json:"DiskName"`
	Description      string `json:"Description"`
	Type             string `json:"Type"`
	Category         string `json:"Category"`
	Size             int    `json:"Size"`
	ImageId          string `json:"ImageId"`
	SourceSnapshotId string `json:"SourceSnapshotId"`
	ProductCode      string `json:"ProductCode"`
	Portable         string `json:"Portable"`
	Status           string `json:"Status"`
	OperationLocks   struct {
		OperationLock OperationLocksType `json:"OperationLock"`
	} `json:"OperationLocks"`
	InstanceId         string `json:"InstanceId"`
	Device             string `json:"Device"`
	DeleteWithInstance string `json:"DeleteWithInstance"`
	DeleteAutoSnapshot string `json:"DeleteAutoSnapshot"`
	EnableAutoSnapshot string `json:"EnableAutoSnapshot"`
	CreationTime       string `json:"CreationTime"`
	AttachedTime       string `json:"AttachedTime"`
	DetachedTime       string `json:"DetachedTime"`
}

type DiskSetType struct {
	Disk DiskItemType `json:"Disks"`
}

//镜像类型
type ImageType struct {
	ImageId            string `json:"ImageId"`
	ImageVersion       string `json:"ImageVersion"`
	Architecture       string `json:"Architecture"`
	ImageName          string `json:"ImageName"`
	Description        string `json:"Description"`
	Size               int64  `json:"Size"`
	ImageOwnerAlias    string `json:"ImageOwnerAlias"`
	OSName             string `json:"OSName"`
	DiskDeviceMappings struct {
		DiskDeviceMapping []DiskDeviceMapping `json:"DiskDeviceMapping"`
	} `json:"DiskDeviceMappings"`
	ProductCode  string    `json:"ProductCode"`
	IsSubscribed bool      `json:"IsSubscribed"`
	Progress     string    `json:"Progress"`
	Status       string    `json:"Status"`
	CreationTime time.Time `json:"CreationTime"`
	IsSelfShared string    `json:"IsSelfShared"`
	IsCopied     bool      `json:"IsCopied"`
	Usage        string    `json:"Usage"`
	Tags         struct {
		Tag []string `json:"Tag"`
	} `json:"Tags"`
	Platform string `json:"Platform "`
}

//实例监控字段
type InstanceMonitorDataType struct {
	InstanceId        string `json:"InstanceId"`
	CPU               int    `json:"CPU"`
	IntranetRX        int    `json:"IntranetRX"`
	IntranetTX        int    `json:"IntranetTX"`
	IntranetBandwidth int    `json:"IntranetBandwidth"`
	InternetRX        int    `json:"InternetRX"`
	InternetTX        int    `json:"InternetTX"`
	InternetBandwidth int    `json:"InternetBandwidth"`
	IOPSRead          int    `json:"IOPSRead"`
	IOPSWrite         int    `json:"IOPSWrite"`
	BPSRead           int    `json:"BPSRead"`
	BPSWrite          int    `json:"BPSWrite"`
	TimeStamp         string `json:"TimeStamp"`
}

//实例状态
type InstanceStatusItemType struct {
	InstanceId string `json:"InstanceId"`
	Status     string `json:"Status"`
}

//实例状态
var INSTANCE_STATUSES = map[string]string{
	"Stopped":  "已停止",
	"Starting": "启动中",
	"Running":  "运行中",
	"Stopping": "停止中",
	"Deleted":  "已释放",
}

//包含实例状态的项的集合
type InstanceStatusSetType struct {
	InstanceStatus []InstanceStatusItemType `json:"InstanceStatus"`
}

//锁定原因
var LOCK_REASONS = map[string]string{
	"financial": "因欠费被锁定",
	"security":  "因安全原因被锁定",
}

//资源的锁定原因类型
type OperationLocksType struct {
	LockReason []string `json:"LockReason"`
}

type InstanceTypeItemType struct {
	InstanceTypeId string  `json:"InstanceTypeId"`
	CpuCoreCount   int     `json:"CpuCoreCount"`
	MemorySize     float64 `json:"MemorySize"`
}

//包含IP地址的集合
type IpAddressSetType struct {
	IpAddress []string `json:"IpAddress"`
}

//安全组规则任命类型
type PermissionSetType struct {
	Permission []PermissionType `json:"Permission"`
}

//安全组规则类型
type PermissionType struct {
	IpProtocol              string `json:"IpProtocol"`
	PortRange               string `json:"PortRange"`
	SourceCidrIp            string `json:"SourceCidrIp"`
	SourceGroupId           string `json:"SourceGroupId"`
	SourceGroupOwnerAccount string `json:"SourceGroupOwnerAccount"`
	DestCidrIp              string `json:"DestCidrIp"`
	DestGroupId             string `json:"DestGroupId"`
	DestGroupOwnerAccount   string `json:"DestGroupOwnerAccount"`
	Policy                  string `json:"Policy"`
	NicType                 string `json:"NicType"`
	Priority                string `json:"Priority"`
	Direction               string `json:"Direction"`
	Description             string `json:"Description"`
}

//Region信息的类型
type RegionType struct {
	RegionId  string `json:"RegionId"`
	LocalName string `json:"LocalName"`
}

//资源类型
var RESOURCE_TYPES = map[string]string{
	"Instance":    "支持实例创建",
	"IoOptimized": "支持IO优化的实例创建",
	"Disk":        "支持硬盘创建",
	"VSwitch":     "支持转悠网络创建",
}

//允许创建的资源类型
type AvailableResourceCreationType struct {
	ResourceTypes []string `json:"ResourceTypes"`
}

//磁盘种类
var DISK_CATEGORIES = map[string]string{
	"cloud":         "支持创建普通云盘和独立普通云盘",
	"cloud_ssd":     "支持创建SSD云盘",
	"ephemeral":     "支持创建本地磁盘",
	"ephemeral_ssd": "支持创建本地 SSD 盘",
}

//支持的磁盘种类
type AvailableDiskCategoriesType struct {
	DiskCategories []string `json:"DiskCategories"`
}

//可用区信息的类型
type ZoneType struct {
	ZoneId                    string                        `json:"ZoneId"`
	LocalName                 string                        `json:"LocalName"`
	AvailableResourceCreation AvailableResourceCreationType `json:"AvailableResourceCreation"`
	AvailableDiskCategories   AvailableDiskCategoriesType   `json:"AvailableDiskCategories"`
}

type ClusterType struct {
	ClusterId string `json:"ClusterType"`
}

type SnapshotType struct {
	SnapshotId     string `json:"SnapshotId"`
	SnapshotName   string `json:"SnapshotName"`
	Description    string `json:"Description"`
	Progress       string `json:"Progress"`
	SourceDiskId   string `json:"SourceDiskId"`
	SourceDiskSize int    `json:"SourceDiskSize"`
	SourceDiskType string `json:"SourceDiskType"`
	ProductCode    string `json:"ProductCode"`
	CreationTime   string `json:"CreationTime"`
}

//安全组ID集合的数据类型
type SecurityGroupIdSetType struct {
	SecurityGroupId []string `json:"SecurityGroupId"`
}

type SecurityGroupSetType struct {
	SecurityGroup SecurityGroupItemType `json:"SecurityGroup"`
}

//安全组集合的类型
type SecurityGroupItemType struct {
	SecurityGroupId   string    `json:"SecurityGroupId"`
	SecurityGroupName string    `json:"SecurityGroupName"`
	Description       string    `json:"Description"`
	VpcId             string    `json:"VpcId"`
	CreationTime      time.Time `json:"CreationTime"`
}

type IpRangeSetType struct {
	IpAddress string `json:"IpAddress"`
	NicType   string `json:"NicType"`
}

var PERIOD_MAP = map[int]string{
	1: "1:00 - 7:00",
	2: "7:00 - 13:00",
	3: "13:00 - 19:00",
	4: "19:00 - 1:00",
}

type AutoSnapshotPolicyType struct {
	SystemDiskPolicyEnabled           string `json:"SystemDiskPolicyEnabled"`
	SystemDiskPolicyTimePeriod        int    `json:"SystemDiskPolicyTimePeriod"`
	SystemDiskPolicyRetentionDays     int    `json:"SystemDiskPolicyRetentionDays"`
	SystemDiskPolicyRetentionLastWeek string `json:"SystemDiskPolicyRetentionLastWeek"`
	DataDiskPolicyEnabled             string `json:"DataDiskPolicyEnabled"`
	DataDiskPolicyTimePeriod          string `json:"DataDiskPolicyTimePeriod"`
	DataDiskPolicyRetentionDays       int    `json:"DataDiskPolicyRetentionDays"`
	DataDiskPolicyRetentionLastWeek   int    `json:"DataDiskPolicyRetentionLastWeek"`
}

type AutoSnapshotExecutionStatusType struct {
	SystemDiskExecutionStatus string `json:"SystemDiskExecutionStatus"`
	DataDiskExecutionStatus   string `json:"DataDiskExecutionStatus"`
}

//磁盘设备
type DiskDeviceMapping struct {
	SnapshotId string `json:"SnapshotId"`
	Size       string `json:"Size"`
	Device     string `json:"Device"`
}

type VpcSetType struct {
	VpcId        string `json:"VpcId"`
	RegionId     string `json:"RegionId"`
	Status       string `json:"Status"`
	VpcName      string `json:"VpcName"`
	VSwitchIds   string `json:"VSwitchIds"`
	CidrBlock    string `json:"CidrBlock"`
	VRouterId    string `json:"VRouterId"`
	Description  string `json:"Description"`
	CreationTime string `json:"CreationTime"`
}

type VRouterSetType struct {
	VRouterId     string `json:"VRouterId"`
	RegionId      string `json:"RegionId"`
	VpcId         string `json:"VpcId"`
	RouteTableIds string `json:"RouteTableIds"`
	VRouterName   string `json:"VRouterName"`
	Description   string `json:"Description"`
	CreationTime  string `json:"CreationTime"`
}

type RouteTableSetType struct {
	VRouterId      string              `json:"VRouterId"`
	RouteTableId   string              `json:"RouteTableId"`
	RouteEntrys    []RouteEntrySetType `json:"RouteEntrys"`
	RouteTableType string              `json:"RouteTableType"`
	CreationTime   string              `json:"CreationTime"`
}

type RouteEntrySetType struct {
	RouteTableId         string `json:"RouteTableId"`
	DestinationCidrBlock string `json:"DestinationCidrBlock"`
	Type                 string `json:"Type"`
	NextHopId            string `json:"NextHopId"`
	Status               string `json:"Status"`
}

type VSwitchSetType struct {
	VSwitchId               string `json:"VSwitchId"`
	VpcId                   string `json:"VpcId"`
	Status                  string `json:"Status"`
	CidrBlock               string `json:"CidrBlock"`
	ZoneId                  string `json:"ZoneId"`
	AvailableIpAddressCount int    `json:"AvailableIpAddressCount"`
	Description             string `json:"Description"`
	VSwitchName             string `json:"VSwitchName"`
	CreationTime            string `json:"CreationTime"`
}

//弹性公网IP绑定信息
type EipAddressAssociateType struct {
	AllocationId       string `json:"AllocationId"`
	IpAddress          string `json:"IpAddress"`
	Bandwidth          int    `json:"Bandwidth"`
	InternetChargeType string `json:"InternetChargeType"`
}

type EipAddressSetType struct {
	RegionId           string             `json:"RegionId"`
	IpAddress          string             `json:"IpAddress"`
	AllocationId       string             `json:"AllocationId"`
	Status             string             `json:"Status"`
	InstanceId         string             `json:"InstanceId"`
	Bandwidth          string             `json:"Bandwidth"`
	InternetChargeType string             `json:"InternetChargeType"`
	OperationLocks     OperationLocksType `json:"OperationLocks"`
	AllocationTime     string             `json:"AllocationTime"`
}

var INTERNET_CHARGE_TYPES = map[string]string{
	"PayByBandwidth": "按固定带宽付费",
	"PayByTraffic":   "按流量付费",
}

//云服务器实例属性
type InstanceAttributesType struct {
	ClusterID               string                  `json:"ClusterId"`
	CreationTime            string                  `json:"CreationTime"`
	Description             string                  `json:"Description"`
	HostName                string                  `json:"HostName"`
	ImageID                 string                  `json:"ImageId"`
	InnerIpAddress          IpAddressSetType        `json:"InnerIpAddress"`
	VlanID                  string                  `json:"VlanId"`
	InstanceId              string                  `json:"InstanceId"`
	InstanceName            string                  `json:"InstanceName"`
	RegionId                string                  `json:"RegionId"`
	ZoneId                  string                  `json:"ZoneId"`
	InstanceType            string                  `json:"InstanceType"`
	Status                  string                  `json:"Status"`
	SecurityGroupIds        SecurityGroupIdSetType  `json:"SecurityGroupIds"`
	PublicIpAddress         IpAddressSetType        `json:"PublicIpAddress"`
	InternetMaxBandwidthIn  int64                   `json:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut int64                   `json:"InternetMaxBandwidthOut"`
	InternetChargeType      string                  `json:"InternetChargeType"`
	VpcAttributes           VpcAttributesType       `json:"VpcAttributes"`
	EipAddress              EipAddressAssociateType `json:"EipAddress"`
	InstanceNetworkType     string                  `json:"InstanceNetworkType"`
	OperationLocks          OperationLocksType      `json:"OperationLocks"`
}

//云服务器实例的专有网络相关属性
type VpcAttributesType struct {
	VpcId            string           `json:"VpcId"`
	VSwitchId        string           `json:"VSwitchId"`
	PrivateIpAddress IpAddressSetType `json:"PrivateIpAddress"`
	NatIpAddress     string           `json:"NatIpAddress"`
}

type EipMonitorDataType struct {
	EipRX        int    `json:"EipRX"`
	EipTX        int    `json:"EipTX"`
	EipFlow      int    `json:"EipFlow"`
	EipBandwidth int    `json:"EipBandwidth"`
	EipPackets   int    `json:"EipPackets"`
	TimeStamp    string `json:"TimeStamp"`
}

type DiskMonitorDataType struct {
	DiskId    string `json:"DiskId"`
	IOPSRead  int    `json:"IOPSRead"`
	IOPSWrite int    `json:"IOPSWrite"`
	IOPSTotal int    `json:"IOPSTotal"`
	BPSRead   int    `json:"BPSRead"`
	BPSWrite  int    `json:"BPSWrite"`
	BPSTotal  int    `json:"BPSTotal"`
	TimeStamp string `json:"TimeStamp"`
}

type AccountType struct {
	AliyunId string `json:"AliyunId"`
}

type ShareGroupType struct {
	Group string `json:"Group"`
}
