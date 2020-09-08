package rds

import (
	//"net"
	"time"
)

var RDS_STATUSES = map[string]string{
	"Creating":                  "创建中",
	"Running":                   "使用中",
	"Deleting":                  "删除中",
	"Rebooting":                 "重启中",
	"DBInstanceClassChanging":   "升降级中",
	"TRANSING":                  "迁移中",
	"EngineVersionUpgrading":    "迁移版本中",
	"TransingToOthers":          "迁移数据到其他RDS中",
	"GuardDBInstanceCreating":   "生产灾备实例中",
	"Restoring":                 "备份恢复中",
	"Importing":                 "数据导入中",
	"ImportingFromOthers":       "从其他RDS实例导入数据中",
	"DBInstanceNetTypeChanging": "内外网切换中",
	"GuardSwitching":            "容灾切换中",
}

const (
	ACCOUNT_PRIVILEGE_READONLY  = "ReadOnly"
	ACCOUNT_PRIVILEGE_READWRITE = "ReadWrite"
)

const (
	DB_STATUS_CREATING = "Creating"
	DB_STATUS_RUNNING  = "Running"
	DB_STATUS_DELETING = "Deleting"
)

const (
	DBINSTANCE_STATUS_CREATING                  = "Creating"
	DBINSTANCE_STATUS_RUNNING                   = "Running"
	DBINSTANCE_STATUS_DELETING                  = "Deleting"
	DBINSTANCE_STATUS_REBOOTING                 = "Rebooting"
	DBINSTANCE_STATUS_DBINSTANCECLASSCHANGING   = "DBInstanceClassChanging"
	DBINSTANCE_STATUS_TRANSING                  = "TRANSING"
	DBINSTANCE_STATUS_ENGINEVERSIONUPGRADING    = "EngineVersionUpgrading"
	DBINSTANCE_STATUS_TRANSINGTOOTHERS          = "TransingToOthers"
	DBINSTANCE_STATUS_GUARDDBINSTANCECREATING   = "TransingToOthers"
	DBINSTANCE_STATUS_RESTORING                 = "TransingToOthers"
	DBINSTANCE_STATUS_IMPORTING                 = "Importing"
	DBINSTANCE_STATUS_IMPORTINGFROMOTHERS       = "ImportingFromOthers"
	DBINSTANCE_STATUS_DBINSTANCENETTYPECHANGING = "DBInstanceNetTypeChanging"
	DBINSTANCE_STATUS_GUARDSWITCHING            = "GuardSwitching"
)

type DBInstance struct {
	ConnectionMode        string  `json:"ConnectionMode"`
	CreateTime            string  `json:"CreateTime"`
	DBInstanceDescription string  `json:"DBInstanceDescription"`
	DBInstanceID          string  `json:"DBInstanceId"`
	DBInstanceNetType     string  `json:"DBInstanceNetType"`
	DBInstanceStatus      string  `json:"DBInstanceStatus"`
	DBInstanceType        string  `json:"DBInstanceType"`
	Engine                string  `json:"Engine"`
	EngineVersion         float64 `json:"EngineVersion,string"`
	ExpireTime            string  `json:"ExpireTime,omitempty"`
	GuardDBInstanceID     string  `json:"GuardDBInstanceId"`
	InstanceNetworkType   string  `json:"InstanceNetworkType"`
	LockMode              string  `json:"LockMode"`
	LockReason            string  `json:"LockReason"`
	MasterInstanceID      string  `json:"MasterInstanceId"`
	MutriORsignle         bool    `json:"MutriORsignle"`
	PayType               string  `json:"PayType"`
	ReadOnlyDBInstanceIDs struct {
		ReadOnlyDBInstanceID []interface{} `json:"ReadOnlyDBInstanceId"`
	} `json:"ReadOnlyDBInstanceIds"`
	RegionID         string `json:"RegionId"`
	TempDBInstanceID string `json:"TempDBInstanceId"`
	VpcID            string `json:"VpcId"`
	ZoneID           string `json:"ZoneId"`
	InsID            int64  `json:"InsId"`
}

type DBInstanceAttribute struct {
	AccountMaxQuantity          int64     `json:"AccountMaxQuantity"`
	AvailabilityValue           string    `json:"AvailabilityValue"`
	ConnectionMode              string    `json:"ConnectionMode"`
	ConnectionString            string    `json:"ConnectionString"`
	CreationTime                time.Time `json:"CreationTime"`
	DBInstanceClass             string    `json:"DBInstanceClass"`
	DBInstanceDescription       string    `json:"DBInstanceDescription"`
	DBInstanceID                string    `json:"DBInstanceId"`
	DBInstanceMemory            int64     `json:"DBInstanceMemory"`
	DBInstanceNetType           string    `json:"DBInstanceNetType"`
	DBInstanceStatus            string    `json:"DBInstanceStatus"`
	DBInstanceStorage           int64     `json:"DBInstanceStorage"`
	DBInstanceType              string    `json:"DBInstanceType"`
	DBMaxQuantity               int64     `json:"DBMaxQuantity"`
	Engine                      string    `json:"Engine"`
	EngineVersion               float64   `json:"EngineVersion,string"`
	ExpireTime                  time.Time `json:"ExpireTime"`
	GuardDBInstanceID           string    `json:"GuardDBInstanceId"`
	IncrementSourceDBInstanceID string    `json:"IncrementSourceDBInstanceId"`
	InstanceNetworkType         string    `json:"InstanceNetworkType"`
	LockMode                    string    `json:"LockMode"`
	LockReason                  string    `json:"LockReason"`
	MaintainTime                string    `json:"MaintainTime"`
	MasterInstanceID            string    `json:"MasterInstanceId"`
	MaxConnections              int64     `json:"MaxConnections"`
	MaxIOPS                     int64     `json:"MaxIOPS"`
	PayType                     string    `json:"PayType"`
	Port                        int64     `json:"Port,string"`
	ReadOnlyDBInstanceIDs       struct {
		ReadOnlyDBInstanceID []interface{} `json:"ReadOnlyDBInstanceId"`
	} `json:"ReadOnlyDBInstanceIds"`
	RegionID         string `json:"RegionId"`
	SecurityIPList   string `json:"SecurityIPList"`
	//SecurityIPList   net.IP `json:"SecurityIPList"`
	TempDBInstanceID string `json:"TempDBInstanceId"`
	ZoneID           string `json:"ZoneId"`
}

type DataBase struct {
	Accounts struct {
		AccountPrivilegeInfo []struct {
			Account          string `json:"Account"`
			AccountPrivilege string `json:"AccountPrivilege"`
		} `json:"AccountPrivilegeInfo"`
	} `json:"Accounts"`
	CharacterSetName string `json:"CharacterSetName"`
	DBDescription    string `json:"DBDescription"`
	DBInstanceID     string `json:"DBInstanceId"`
	DBName           string `json:"DBName"`
	DBStatus         string `json:"DBStatus"`
	Engine           string `json:"Engine"`
}

type DBInstanceAccount struct {
	AccountDescription string `json:"AccountDescription"`
	AccountName        string `json:"AccountName"`
	AccountStatus      string `json:"AccountStatus"`
	DBInstanceID       string `json:"DBInstanceId"`
	DbInstanceName     string `json:"DbInstanceName"`
	DatabasePrivileges struct {
		DatabasePrivilege []struct {
			DBName           string `json:"DBName"`
			AccountPrivilege string `json:"AccountPrivilege"`
		} `json:"DatabasePrivilege"`
	} `json:"DatabasePrivileges"`
}
