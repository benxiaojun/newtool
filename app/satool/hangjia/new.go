package main

import (
	"fmt"
	"github.com/benxiaojun/satool/aliyun/rds"
	"github.com/benxiaojun/satool/app/satool"
	"github.com/benxiaojun/satool/qiniu"
	"time"
)

func newsite(site, sitename, region, ecsid, rdsid, rdsname, rdspassword, rdsurmname, rdsurmpassword, qiniuprefix, qiniuname, qiniupassword string) {
	//更新ECS备注
	satool.Modifyecs(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, ecsid, sitename)
	//获取ecs信息
	ecs, err := satool.Getecs(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, ecsid, region)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		ecspublicip := ecs.PublicIpAddress.IpAddress[0]
		ecsinternalip := ecs.InnerIpAddress.IpAddress[0]
		fmt.Printf("ECS外网IP: %s\nECS内网IP: %s\n", ecspublicip, ecsinternalip)

		//更新RDS备注
		satool.Modifyrds(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, sitename)
		//获取rds信息
		rdsinfo, err := satool.Getrds(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			rdsurl := rdsinfo.ConnectionString
			fmt.Printf("RDS地址: %s\n", rdsurl)
			fmt.Printf("创建%s相关...\n", site)
			//创建site rds数据库
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsname, sitename); ok {
				fmt.Printf("创建%s数据库成功: %s\n", site, rdsname)
				//创建rds账号
				if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsname, rdspassword, sitename); ok {
					fmt.Printf("创建%s账号成功: %s[%s]\n", site, rdsname, rdspassword)
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
							fmt.Print("FAMILY数据库创建中，等待3s再次检查...\n")
							time.Sleep(3 * time.Second)
						case rds.DB_STATUS_DELETING:
							fmt.Print("FAMILY数据库删除中...\n")
							break
						}
					}
					accreated := false
					for !accreated {
						acstatus, err := satool.Getaccountstatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsname)
						if err != nil {
							fmt.Printf("%v\n", err)
							break
						}
						switch acstatus {
						case "Unavailable":
							fmt.Print("FAMILY账号不可用，等待3s再次检查...\n")
							time.Sleep(3 * time.Second)
							break
						case "Available":
							accreated = true
						}
					}
					if dbcreated && accreated {
						//授权
						if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsname, rdsname); ok {
							fmt.Print("FAMILY授权成功\n")

						} else {
							fmt.Printf("%v\n", err)
						}
					}
				} else {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Printf("%v\n", err)
			}
			fmt.Println("创建URM相关...")
			//创建urm rds数据库
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsurmname, sitename); ok {
				fmt.Printf("创建URM数据库成功: %s\n", rdsurmname)
				//创建rds账号
				if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsurmname, rdsurmpassword, sitename); ok {
					fmt.Printf("创建URM账号成功: %s[%s]\n", rdsurmname, rdsurmpassword)
					//检查数据库状态
					dbcreated := false
					for !dbcreated {
						dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsurmname)
						if err != nil {
							fmt.Printf("%v\n", err)
							break
						}
						switch dbstatus {
						case rds.DB_STATUS_RUNNING:
							dbcreated = true
						case rds.DB_STATUS_CREATING:
							fmt.Print("URM数据库创建中，等待3s再次检查...\n")
							time.Sleep(3 * time.Second)
						case rds.DB_STATUS_DELETING:
							fmt.Print("URM数据库删除中...\n")
							break
						}
					}
					accreated := false
					for !accreated {
						acstatus, err := satool.Getaccountstatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsurmname)
						if err != nil {
							fmt.Printf("%v\n", err)
							break
						}
						switch acstatus {
						case "Unavailable":
							fmt.Print("URM账号不可用，等待3s再次检查...\n")
							time.Sleep(3 * time.Second)
							break
						case "Available":
							accreated = true
						}
					}
					if dbcreated && accreated {
						//授权
						if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsurmname, rdsurmname); ok {
							fmt.Print("URM授权成功\n")

						} else {
							fmt.Printf("%v\n", err)
						}
					}
				} else {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Printf("%v\n", err)
			}
			//设置白名单
			if ok, err := satool.Modifysecurityips(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, ecsinternalip, rdsname, ""); ok {
				fmt.Printf("设置数据库白名单成功: %s\n", ecsinternalip)
				//是否升级数据库
				goon := true
				if rdsinfo.EngineVersion < 5.6 {
					fmt.Printf("数据库版本低，升级数据库\n")
					if ok, err := satool.Upgraderdsengine(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid); !ok {
						fmt.Printf("%v\n", err)
						goon = false
					}
				}
				if goon {
					if len(qiniuname) > 0 {
						access, err := satool.Getqiniuaccess(QINIU_USERNAME, QINIU_PASSWORD)
						if err != nil {
							fmt.Printf("%v\n", err)
						} else {
							fmt.Print("获取七牛授权成功\n")
							qiniuemail := QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX
							fmt.Printf("七牛子账号: %s[%s]\n", qiniuemail, qiniupassword)
							//创建七牛子账号
							uid, _ := satool.Getqiniuuid(*access, qiniuemail)
							var child qiniu.Child
							if uid == 0 {
								qiniu, err := satool.Createqiniuaccount(*access, QINIU_EMAIL_PREFIX+qiniuname+QINIU_EMAIL_SUFFIX, qiniupassword)
								if err != nil {
									fmt.Printf("%v\n", err)
								} else {
									fmt.Printf("子账号创建成功: %d\n", qiniu.UID)
									uid = qiniu.UID
									info, err := satool.Getqiniukey(*access, qiniu.UID)
									if err != nil {
										fmt.Printf("%v\n", err)
									} else {
										qiniuak := info.Key
										qiniusk := info.Secret
										fmt.Printf("AK: %s\nSK: %s\n", qiniuak, qiniusk)
										child, _ = satool.Getqiniuchild(qiniu.UID, *access)
									}
								}
							} else {
								fmt.Printf("找到已创建的子账号: %d\n", uid)
								child, _ = satool.Getqiniuchild(uid, *access)
							}
							if len(child.Key) > 0 && len(child.Secret) > 0 {
								//创建bucket
								if ok, err := satool.Createqiniubucket(child, uid, qiniuprefix+qiniuname); ok {
									fmt.Printf("创建子账号数据空间[%s]成功\n", qiniuprefix+qiniuname)
									url, err := satool.Getqiniubucket(child, qiniuprefix+qiniuname)
									if err != nil {
										fmt.Printf("%v\n", err)
									} else {
										qiniuurl := url
										fmt.Printf("数据空间链接: %s\n", qiniuurl)

										fmt.Print("操作完成\n")
									}

								} else {
									fmt.Printf("%v\n", err)
								}
							} else {
								fmt.Println("未能获取子账号授权")
							}
						}
					} else {
						fmt.Print("跳过七牛操作\n")
					}
				}
			} else {
				fmt.Printf("%v\n", err)
			}
		}
	}
}
