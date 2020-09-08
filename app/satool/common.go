package satool

import (
	"encoding/json"
	"fmt"
	"github.com/benxiaojun/satool/aliyun/ecs"
	"github.com/benxiaojun/satool/aliyun/rds"
	"github.com/benxiaojun/satool/dnspod"
	"github.com/benxiaojun/satool/qiniu"
	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gohttp"
	"os"
)

//计算日流量
func Transfer() {
	var inavg, outavg float64
	fmt.Print("平均每秒IN(bit/s):")
	fmt.Scanf("%f\n", &inavg)

	fmt.Print("平均每秒OUT(bit/s):")
	fmt.Scanf("%f\n", &outavg)

	secondtotal := inavg + outavg
	daytotal := secondtotal * 60 * 60 * 24 / 1000 / 1000 * 125 / 1024 / 1024
	fmt.Printf("平均日总流量(GB): %.4f\n", daytotal)
}

//获取ECS信息
func Getecs(id, secret, ecsid, region string) (ecs.InstanceAttributesType, error) {
	req := &ecs.DescribeInstancesRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.RegionId = region
	bytes, _ := json.Marshal([]string{ecsid})
	req.InstanceIds = string(bytes)

	resp := &ecs.DescribeInstancesResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return ecs.InstanceAttributesType{}, err
	} else {
		if resp.OK() {
			return resp.Instances.Instance[0], nil
		} else {
			return ecs.InstanceAttributesType{}, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//更新ECS备注
func Modifyecs(id, secret, ecsid, name string) (bool, error) {
	req := &ecs.ModifyInstanceAttributeRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.InstanceId = ecsid
	req.InstanceName = name
	resp := &ecs.ModifyInstanceAttributeResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//更新RDS备注
func Modifyrds(id, secret, rdsid, name string) (bool, error) {
	req := &rds.ModifyDBInstanceDescriptionRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.DBInstanceDescription = name
	resp := &rds.ModifyDBDescriptionResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//获取RDS信息
func Getrds(id, secret, rdsid string) (rds.DBInstanceAttribute, error) {
	req := &rds.DescribeDBInstanceAttributeRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid

	resp := &rds.DescribeDBInstanceAttributeResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return rds.DBInstanceAttribute{}, err
	} else {
		if resp.OK() {
			return resp.Items.DBInstanceAttribute[0], nil
		} else {
			return rds.DBInstanceAttribute{}, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//获取数据库状态
func Getdatabasestatus(id, secret, rdsid, name string) (string, error) {
	req := &rds.DescribeDatabasesRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.DBName = name

	resp := &rds.DescribeDatabasesResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return "", err
	} else {
		if resp.OK() {
			if len(resp.Databases.DataBase) > 0 {
				return resp.Databases.DataBase[0].DBStatus, nil
			} else {
				return "", fmt.Errorf("数据库未创建")
			}
		} else {
			return "", fmt.Errorf("%s", resp.Msg())
		}
	}
}

//获取账号状态
func Getaccountstatus(id, secret, rdsid, name string) (string, error) {
	req := &rds.DescribeAccountsRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.AccountName = name

	resp := &rds.DescribeAccountsResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return "", err
	} else {
		if resp.OK() {
			if len(resp.Accounts.DBInstanceAccount) > 0 {
				return resp.Accounts.DBInstanceAccount[0].AccountStatus, nil
			} else {
				return "", fmt.Errorf("账号未创建")
			}
		} else {
			return "", fmt.Errorf("%s", resp.Msg())
		}
	}
}

//修改白名单
func Modifysecurityips(id, secret, rdsid, ip, name, bute string) (bool, error) {
	req := &rds.ModifySecurityIpsRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.SecurityIps = ip
	req.DBInstanceIPArrayName = name
	req.DBInstanceIPArrayAttribute = bute
	//req.InstanceName = name
	resp := &rds.ModifySecurityIpsResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//创建RDS账号
func Createrdsaccount(id, secret, rdsid, name, password, sitename string) (bool, error) {
	req := &rds.CreateAccountRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.AccountName = name
	req.AccountPassword = password
	req.AccountDescription = sitename
	resp := &rds.CreateAccountResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//创建数据库
func Createrdsdatabase(id, secret, rdsid, name, sitename string) (bool, error) {
	req := &rds.CreateDatabaseRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.DBName = name
	req.CharacterSetName = "utf8mb4"
	req.DBDescription = sitename
	resp := &rds.CreateDatabaseResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//授权账号数据库
func Grantrdsprivilege(id, secret, rdsid, dbname, username string) (bool, error) {
	req := &rds.GrantAccountPrivilegeRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.AccountName = username
	req.DBName = dbname
	req.AccountPrivilege = rds.ACCOUNT_PRIVILEGE_READWRITE
	resp := &rds.GrantAccountPrivilegeResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//升级数据库引擎
func Upgraderdsengine(id, secret, rdsid string) (bool, error) {
	req := &rds.UpgradeDBInstanceEngineVersionRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.DBInstanceId = rdsid
	req.EngineVersion = "5.6"
	resp := &rds.UpgradeDBInstanceEngineVersionResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

func Getqiniuaccess(username, password string) (*qiniu.Access, error) {
	access := &qiniu.Access{}
	access.Username = username
	access.Password = password
	if gofile.FileExists(qiniu.GetDir() + string(os.PathSeparator) + access.GetFileName()) {
		err := access.ReadOauth(true)
		if err != nil {
			return nil, err
		}
	} else {
		err := access.Get()
		if err != nil {
			return nil, err
		}
	}
	return access, nil
}

func Getqiniuchild(uid int64, access qiniu.Access) (qiniu.Child, error) {
	child := qiniu.Child{}
	child.Uid = uid
	path := qiniu.GetDir() + string(os.PathSeparator) + child.GetFileName()
	if gofile.FileExists(path) {
		err := child.ReadKeys(true, access.AccessToken)
		if err != nil {
			return qiniu.Child{}, err
		}
	}
	if len(child.Key) == 0 || len(child.Secret) == 0 || child.IsExpired() {
		err := child.Get(access.AccessToken)
		if err != nil {
			return qiniu.Child{}, err
		}
	}
	return child, nil
}

func Getqiniukey(access qiniu.Access, uid int64) (qiniu.ChildKeysResponse, error) {
	req := &qiniu.ChildKeysRequest{}
	req.AuthorizationValue = access.AccessToken
	req.Uid = uid
	resp := &qiniu.ChildKeysResponse{}
	err := gohttp.DoPostResponse(req, resp)
	if err != nil {
		return qiniu.ChildKeysResponse{}, err
	} else {
		if resp.OK() {
			return *resp, nil
		} else {
			return qiniu.ChildKeysResponse{}, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//创建七牛子账号
func Createqiniuaccount(access qiniu.Access, email, password string) (qiniu.CreateChildAccountResponse, error) {
	req := &qiniu.CreateChildAccountRequest{}
	req.AuthorizationValue = access.AccessToken
	req.Email = email
	req.Password = password
	resp := &qiniu.CreateChildAccountResponse{}
	err := gohttp.DoPostResponse(req, resp)
	if err != nil {
		return qiniu.CreateChildAccountResponse{}, err
	} else {
		if resp.OK() {
			return *resp, nil
		} else {
			return qiniu.CreateChildAccountResponse{}, fmt.Errorf("%s", resp.Msg())
		}
	}
}

//创建七牛数据空间
func Createqiniubucket(child qiniu.Child, uid int64, bucket string) (bool, error) {
	req := &qiniu.KeyCreateBucketRequest{}
	req.AccessKey = child.Key
	req.SecretKey = child.Secret
	req.Name = bucket
	req.IsPublic = qiniu.BUCKET_PUBLIC
	resp := &qiniu.KeyCreateBucketResponse{}
	err := gohttp.DoPostResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}

func Getqiniubucket(child qiniu.Child, bucket string) (string, error) {
	req := &qiniu.KeyListBucketDomainRequest{}
	req.BucketName = bucket
	req.AccessKey = child.Key
	req.SecretKey = child.Secret
	bytes, err := gohttp.DoGet(req)
	if err != nil {
		return "", err
	} else {
		if string(bytes[0]) == "[" {
			//array
			resp := make([]string, 0)
			err := json.Unmarshal(bytes, &resp)
			if err != nil {
				return "", err
			} else {
				return fmt.Sprintf("%v", resp), nil
			}
		} else {
			//object
			resp := &qiniu.KeyListBucketDomainResponse{}
			err := json.Unmarshal(bytes, resp)
			if err != nil {
				return "", err
			} else {
				return "", fmt.Errorf("%s", resp.Msg())
			}
		}
	}
}

func Getqiniuuid(access qiniu.Access, email string) (int64, error) {
	req := &qiniu.ListChildAccountRequest{}
	req.AuthorizationValue = access.AccessToken
	req.Offset = 0
	req.Limit = 1000
	bytes, err := gohttp.DoPost(req)
	if err != nil {
		return 0, err
	} else {
		if string(bytes[0]) == "[" {
			//array
			resp := []qiniu.ListChildAccountResponse{}
			err := json.Unmarshal(bytes, &resp)
			if err != nil {
				return 0, err
			} else {
				for _, item := range resp {
					if item.Email == email {
						return item.Uid, nil
					}
				}
				return 0, fmt.Errorf("未找到此账号: %s", email)
			}
		} else {
			//object
			resp := &qiniu.ListChildAccountResponse{}
			err := json.Unmarshal(bytes, resp)
			if err != nil {
				return 0, err
			} else {
				return 0, fmt.Errorf(resp.Msg())
			}
		}
	}
}

//创建DNSPOD账号
func Creatednspod(token, domain_id, name, ip string) (bool, error) {
	req := &dnspod.RecordCreateRequest{}
	req.LoginToken = token
	req.DomainId = domain_id
	req.SubDomain = name
	req.RecordType = "A"
	req.RecordLine = "默认"
	req.Value = ip
	req.Status = "enable"
	resp := &dnspod.RecordCreateResponse{}
	err := gohttp.DoPostResponse(req, resp)
	if err != nil {
		return false, err
	} else {
		if resp.OK() {
			return true, nil
		} else {
			return false, fmt.Errorf("%s", resp.Msg())
		}
	}
}
