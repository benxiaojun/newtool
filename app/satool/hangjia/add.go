package main

import (
	"fmt"
	"github.com/benxiaojun/satool/aliyun/rds"
	"github.com/benxiaojun/satool/app/satool"
	"time"
)

func addsite(site, sitename, rdsid, rdsname,rdsusername,useexistuser, rdspassword, qiniuprefix, qiniuname string) {
	fmt.Print("正在操作RDS...\n")
	if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsname, sitename); ok {
		fmt.Printf("创建%s数据库成功: %s\n", site, rdsname)
		if useexistuser == "no" {
			//创建rds账号
			if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername, rdspassword, sitename); ok {
				fmt.Printf("创建账号成功: %s[%s]\n", rdsusername, rdspassword)
			}else {
				fmt.Printf("%v\n", err)
			}
		}
		//检查数据库状态
		dbcreated := false
		for !dbcreated {
			dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsname)
			if err != nil {
				fmt.Printf("%v\n", err)
				break
			}
			switch dbstatus {
			case rds.DB_STATUS_RUNNING:
				dbcreated = true
			case rds.DB_STATUS_CREATING:
				fmt.Print("数据库创建中，等待3s再次检查...\n")
				time.Sleep(3 * time.Second)
			case rds.DB_STATUS_DELETING:
				fmt.Print("数据库删除中...\n")
				break
			}
		}
		//检查账号状态
		accreated := false
		for !accreated {
			acstatus, err := satool.Getaccountstatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername)
			if err != nil {
				fmt.Printf("%v\n", err)
				break
			}
			switch acstatus {
			case "Unavailable":
				fmt.Print("账号不可用，等待3s再次检查...\n")
				time.Sleep(3 * time.Second)
				break
			case "Available":
				accreated = true
			}
		}
		if dbcreated && accreated {
			if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsname, rdsusername); ok {
				fmt.Print("授权成功\n")

			} else {
				fmt.Printf("%v\n", err)
			}
		}
	}else {
		fmt.Printf("%v\n", err)
	}

	fmt.Print("正在操作七牛...\n")
	access, err := satool.Getqiniuaccess(QINIU_USERNAME, QINIU_PASSWORD)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Print("获取七牛授权成功\n")
		qiniuemail := QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX

		uid, err := satool.Getqiniuuid(*access, qiniuemail)
		if err != nil && uid > 0 {
			fmt.Printf("%v\n", err)
		} else {
			child, err := satool.Getqiniuchild(uid, *access)
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Print("获取子账号授权成功\n")
				//创建bucket
				if ok, err := satool.Createqiniubucket(child, uid, qiniuprefix+qiniuname); ok {
					fmt.Printf("创建子账号数据空间[%s]成功\n", qiniuprefix+qiniuname)
					url, err := satool.Getqiniubucket(child, qiniuprefix+qiniuname)
					if err != nil {
						fmt.Printf("%v\n", err)
					} else {
						fmt.Printf("数据空间链接: %s\n", url)

						fmt.Print("操作完成\n")
					}

				} else {
					fmt.Printf("%v\n", err)
				}
			}
		}
	}
}
