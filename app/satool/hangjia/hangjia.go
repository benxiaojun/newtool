package main

import (
	"fmt"
	"github.com/benxiaojun/satool/aliyun/rds"
	"github.com/benxiaojun/satool/app/satool"
	"github.com/benxiaojun/satool/qiniu"
	"github.com/sillydong/goczd/godata"
	"github.com/sillydong/goczd/gohttp"
	"os"
	"time"
)

const DEBUG = false
const VERSION = "v1.0 20150725"

const (
	QINIU_USERNAME          = ""
	QINIU_PASSWORD          = ""
	QINIU_EMAIL_PREFIX      = "client_"
	QINIU_EMAIL_SUFFIX      = "@hangjiayun.com"
	QINIU_FAMILY_PREFIX     = "family-"
	QINIU_HOUSE_PREFIX      = "house-"
	QINIU_WUZHEKA_PREFIX    = "ticket-"
	QINIU_MARRY_PREFIX      = "marry-"
	QINIU_BABY_PREFIX       = "baby-"
	QINIU_URM_PREFIX       = "urm-"
	ALIYUN_ACCESSKEY_ID     = ""
	ALIYUN_ACCESSKEY_SECRET = ""
)

func main() {
	var act string
	fmt.Printf("Welcome to SATOOL %s For Hangjia\nType h to view available commands\n", VERSION)

	gohttp.REQUEST_DEBUG = DEBUG

	for {
		fmt.Scanf("%s\n", &act)
		if len(act) > 0 {
			switch act {
			case "q", "exit", "quit":
				fmt.Println("See you...")
				os.Exit(0)
			//case "v", "version":
			//	fmt.Println(VERSION)
			case "h", "help":
				help()
			case "whiteip":
				whiteip()
			case "newfamily":
				newfamily()
			case "addfamily":
				addfamily()
			case "newhouse":
				newhouse()
			case "addhouse":
				addhouse()
			case "shareticket":
				shareticket()
			case "sharehouse":
				sharehouse()
			case "sharefamily":
				sharefamily()
			case "addwuzheka":
				addwuzheka()
			case "newmarry":
				newmarry()
			case "addmarry":
				addmarry()
			case "yuqing":
				yuqing()
			case "newbaby":
				newbaby()
			case "addbaby":
				addbaby()
			case "t":
				satool.Transfer()
			}
		} else {
			help()
		}
		act = ""
	}
}

//显示可用命令信息
func help() {
	fmt.Println(`操作列表:
   // v|version                     打印程序版本
    h|help                        打印帮助信息
    q|exit|quit                   退出程序
    
    newfamily                     新装家装站点
    addfamily                     补充家装站点
    newhouse                      新装房产站点
    addhouse                      补充房产站点
    shareticket                   共享美食站点
    sharehouse                    共享房产站点
    sharefamily                   共享家居站点
    addticket                     补充美食站点
    newmarry                      新装婚嫁站点
    addmarry                      补充婚嫁站点
    yuqing                        开通舆情数据库
    newbaby                       新装亲子站点
    addbaby                       补充亲子站点
    t                             计算日流量
    whiteip                       RDS加白ECS
`)
}


//加白ecsip
func whiteip() {
	var rdsid string
	var ecsip string
	var sitename string

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入ECS IP:")
	fmt.Scanln(&ecsip)

	fmt.Print("输入站点名称:")
	fmt.Scanln(&sitename)

	rdsinfo, err := satool.Getrds(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		if ok, err := satool.Modifysecurityips(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, ecsip, sitename, ""); ok {
			fmt.Printf("设置数据库白名单成功: %s\n", ecsip)
			//是否升级数据库
			//goon := true
			if rdsinfo.EngineVersion < 5.6 {
				fmt.Printf("数据库版本低，升级数据库\n")
				if ok, err := satool.Upgraderdsengine(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid); !ok {
					fmt.Printf("%v\n", err)
				}
			}

		} else {
			fmt.Printf("%v\n", err)
		}
	}
}



//新装家装站点
func newfamily() {
	var sitename string
	var region string
	var ecsid string
	var rdsid, rdsfamilyname, rdsfamilypassword, rdsurmname, rdsurmpassword string
	var qiniuname, qiniupassword string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入ECS ID:")
	fmt.Scanln(&ecsid)

	fmt.Printf("输入ECS地域[cn-hangzhou]:")
	fmt.Scanln(&region)
	if len(region) == 0 {
		region = "cn-hangzhou"
	}

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	rdsfamilyname = "hangjia_family"
	rdsfamilypassword = godata.RandomString(12, godata.ALPHANUMERIC)
	rdsurmname = "hangjia"
	rdsurmpassword = godata.RandomString(12, godata.ALPHANUMERIC)

	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Print("正在操作...\n")

	newsite("FAMILY", sitename, region, ecsid, rdsid, rdsfamilyname, rdsfamilypassword, rdsurmname, rdsurmpassword, QINIU_FAMILY_PREFIX, qiniuname, qiniupassword)
}

//补充家装站点
func addfamily() {
	//创建数据库
	//创建七牛空间
	var sitename string
	var rdsid, rdsdbname, rdspassword string
	var qiniuname string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛名(不含前后缀):")
	fmt.Scanln(&qiniuname)

	rdsdbname = "hangjia_family"
	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)

	addsite("FAMILY", sitename, rdsid, rdsdbname, rdspassword, QINIU_FAMILY_PREFIX, qiniuname)
}

//新装房产站点
func newhouse() {
	var sitename string
	var region string
	var ecsid string
	var rdsid, rdshousename, rdshousepassword, rdsurmname, rdsurmpassword string
	var qiniuname, qiniupassword string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入ECS ID:")
	fmt.Scanln(&ecsid)

	fmt.Printf("输入ECS地域[cn-hangzhou]:")
	fmt.Scanln(&region)
	if len(region) == 0 {
		region = "cn-hangzhou"
	}

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	rdshousename = "hangjia_house"
	rdshousepassword = godata.RandomString(12, godata.ALPHANUMERIC)
	rdsurmname = "hangjia"
	rdsurmpassword = godata.RandomString(12, godata.ALPHANUMERIC)

	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Print("正在操作...\n")

	newsite("HOUSE", sitename, region, ecsid, rdsid, rdshousename, rdshousepassword, rdsurmname, rdsurmpassword, QINIU_HOUSE_PREFIX, qiniuname, qiniupassword)
}

//补充房产站点
func addhouse() {
	//创建数据库
	//创建七牛空间
	var sitename string
	var rdsid, rdsdbname, rdspassword string
	var qiniuname string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛名(不含前后缀):")
	fmt.Scanln(&qiniuname)

	rdsdbname = "hangjia_house"
	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)

	addsite("HOUSE", sitename, rdsid, rdsdbname, rdspassword, QINIU_HOUSE_PREFIX, qiniuname)
}

//新装婚嫁站点
func newmarry() {
	var sitename string
	var region string
	var ecsid string
	var rdsid, rdsmarryname, rdsmarrypassword, rdsurmname, rdsurmpassword string
	var qiniuname, qiniupassword string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入ECS ID:")
	fmt.Scanln(&ecsid)

	fmt.Printf("输入ECS地域[cn-hangzhou]:")
	fmt.Scanln(&region)
	if len(region) == 0 {
		region = "cn-hangzhou"
	}

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	rdsmarryname = "hangjia_marry"
	rdsmarrypassword = godata.RandomString(12, godata.ALPHANUMERIC)
	rdsurmname = "hangjia"
	rdsurmpassword = godata.RandomString(12, godata.ALPHANUMERIC)

	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Print("正在操作...\n")

	newsite("MARRY", sitename, region, ecsid, rdsid, rdsmarryname, rdsmarrypassword, rdsurmname, rdsurmpassword, QINIU_MARRY_PREFIX, qiniuname, qiniupassword)
}

//补充婚嫁站点
func addmarry() {
	//创建数据库
	//创建七牛空间
	var sitename string
	var rdsid, rdsdbname, rdspassword string
	var qiniuname string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛名(不含前后缀):")
	fmt.Scanln(&qiniuname)

	rdsdbname = "hangjia_marry"
	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)

	addsite("MARRY", sitename, rdsid, rdsdbname, rdspassword, QINIU_MARRY_PREFIX, qiniuname)

}

//新装亲子站点
func newbaby() {
	var sitename string
	var region string
	var ecsid string
	var rdsid, rdsbabyname, rdsbabypassword, rdsurmname, rdsurmpassword string
	var qiniuname, qiniupassword string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入ECS ID:")
	fmt.Scanln(&ecsid)

	fmt.Printf("输入ECS地域[cn-hangzhou]:")
	fmt.Scanln(&region)
	if len(region) == 0 {
		region = "cn-hangzhou"
	}

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	rdsbabyname = "hangjia_baby"
	rdsbabypassword = godata.RandomString(12, godata.ALPHANUMERIC)
	rdsurmname = "hangjia"
	rdsurmpassword = godata.RandomString(12, godata.ALPHANUMERIC)

	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Print("正在操作...\n")

	newsite("BABY", sitename, region, ecsid, rdsid, rdsbabyname, rdsbabypassword, rdsurmname, rdsurmpassword, QINIU_BABY_PREFIX, qiniuname, qiniupassword)
}

//补充亲子站点
func addbaby() {
	//创建数据库
	//创建七牛空间
	var sitename string
	var rdsid, rdsdbname, rdspassword string
	var qiniuname string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛名(不含前后缀):")
	fmt.Scanln(&qiniuname)

	rdsdbname = "hangjia_baby"
	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)

	addsite("BABY", sitename, rdsid, rdsdbname, rdspassword, QINIU_BABY_PREFIX, qiniuname)
}

//共享五折卡站点
func shareticket() {
	//创建RDS账号
	//创建RDS库
	//创建RDS账号
	//创建RDS库
	//创建七牛账号
	var sitename string
	var rdsid, rdsusername, rdspassword, wuzhekadbname, urmdbname string
	var qiniuname, qiniupassword, qiniuemail, qiniuak, qiniusk, qiniuurl string
	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入RDS账号名:")
	fmt.Scanln(&rdsusername)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)
	wuzhekadbname = rdsusername + "-ticket"
	urmdbname = rdsusername + "-urm"

	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)
	qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX

	fmt.Print("正在操作RDS...\n")
	//创建rds账号
	if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername, rdspassword, sitename); ok {
		fmt.Printf("创建账号成功: %s[%s]\n", rdsusername, rdspassword)
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
		if accreated {
			//WUZHEKA
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname, sitename); ok {
				fmt.Printf("创建WUZHEKA数据库成功: %s\n", wuzhekadbname)
				dbcreated := false
				for !dbcreated {
					dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname)
					if err != nil {
						fmt.Printf("%v\n", err)
						break
					}
					switch dbstatus {
					case rds.DB_STATUS_RUNNING:
						dbcreated = true
					case rds.DB_STATUS_CREATING:
						fmt.Print("WUZHEKA数据库创建中，等待3s再次检查...\n")
						time.Sleep(3 * time.Second)
					case rds.DB_STATUS_DELETING:
						fmt.Print("WUZHEKA数据库删除中...\n")
						break
					}
				}
				if dbcreated {
					if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname, rdsusername); ok {
						fmt.Print("WUZHEKA授权成功\n")

					} else {
						fmt.Printf("%v\n", err)
					}
				}
			} else {
				fmt.Printf("%v\n", err)
			}

			//URM
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname, sitename); ok {
				fmt.Printf("创建URM数据库成功: %s\n", urmdbname)
				dbcreated := false
				for !dbcreated {
					dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname)
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
				if dbcreated {
					if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname, rdsusername); ok {
						fmt.Print("URM授权成功\n")

					} else {
						fmt.Printf("%v\n", err)
					}
				}
			} else {
				fmt.Printf("%v\n", err)
			}
		}

	} else {
		fmt.Printf("%v\n", err)
	}

	if len(qiniuname) > 0 {
		fmt.Print("正在操作七牛...\n")
		access, err := satool.Getqiniuaccess(QINIU_USERNAME, QINIU_PASSWORD)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("获取七牛授权成功\n")
			qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX
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
						qiniuak = info.Key
						qiniusk = info.Secret
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
				if ok, err := satool.Createqiniubucket(child, uid, QINIU_WUZHEKA_PREFIX+qiniuname); ok {
					fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_WUZHEKA_PREFIX+qiniuname)
					url, err := satool.Getqiniubucket(child, QINIU_WUZHEKA_PREFIX+qiniuname)
					if err != nil {
						fmt.Printf("%v\n", err)
					} else {
						qiniuurl = url
						fmt.Printf("数据空间链接: %s\n", qiniuurl)

						fmt.Print("操作完成\n")
					}
				} else {
					fmt.Printf("%v\n", err)
				}
				if ok, err := satool.Createqiniubucket(child, uid,  QINIU_URM_PREFIX +qiniuname); ok {
					fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_URM_PREFIX  +qiniuname)
					url, err := satool.Getqiniubucket(child, QINIU_URM_PREFIX  +qiniuname)
					if err != nil {
						fmt.Printf("%v\n", err)
					} else {
						qiniuurl = url
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


//共享五折卡站点
func sharehouse() {
	//创建RDS账号
	//创建RDS库
	//创建RDS账号
	//创建RDS库
	//创建七牛账号
	var sitename string
	var rdsid, rdsusername, rdspassword, wuzhekadbname, urmdbname string
	var qiniuname, qiniupassword, qiniuemail, qiniuak, qiniusk, qiniuurl string
	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入RDS账号名:")
	fmt.Scanln(&rdsusername)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)
	wuzhekadbname = rdsusername + "-house"
	urmdbname = rdsusername + "-urm"

	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)
	qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX

	fmt.Print("正在操作RDS...\n")
	//创建rds账号
	if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername, rdspassword, sitename); ok {
		fmt.Printf("创建账号成功: %s[%s]\n", rdsusername, rdspassword)
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
		if accreated {
			//WUZHEKA
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname, sitename); ok {
				fmt.Printf("创建house数据库成功: %s\n", wuzhekadbname)
				dbcreated := false
				for !dbcreated {
					dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname)
					if err != nil {
						fmt.Printf("%v\n", err)
						break
					}
					switch dbstatus {
					case rds.DB_STATUS_RUNNING:
						dbcreated = true
					case rds.DB_STATUS_CREATING:
						fmt.Print("house数据库创建中，等待3s再次检查...\n")
						time.Sleep(3 * time.Second)
					case rds.DB_STATUS_DELETING:
						fmt.Print("house数据库删除中...\n")
						break
					}
				}
				if dbcreated {
					if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname, rdsusername); ok {
						fmt.Print("house授权成功\n")

					} else {
						fmt.Printf("%v\n", err)
					}
				}
			} else {
				fmt.Printf("%v\n", err)
			}

			//URM
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname, sitename); ok {
				fmt.Printf("创建URM数据库成功: %s\n", urmdbname)
				dbcreated := false
				for !dbcreated {
					dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname)
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
				if dbcreated {
					if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname, rdsusername); ok {
						fmt.Print("URM授权成功\n")

					} else {
						fmt.Printf("%v\n", err)
					}
				}
			} else {
				fmt.Printf("%v\n", err)
			}
		}

	} else {
		fmt.Printf("%v\n", err)
	}

	if len(qiniuname) > 0 {
		fmt.Print("正在操作七牛...\n")
		access, err := satool.Getqiniuaccess(QINIU_USERNAME, QINIU_PASSWORD)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("获取七牛授权成功\n")
			qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX
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
						qiniuak = info.Key
						qiniusk = info.Secret
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
				if ok, err := satool.Createqiniubucket(child, uid, QINIU_HOUSE_PREFIX +qiniuname); ok {
					fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_HOUSE_PREFIX +qiniuname)
					url, err := satool.Getqiniubucket(child, QINIU_HOUSE_PREFIX +qiniuname)
					if err != nil {
						fmt.Printf("%v\n", err)
					} else {
						qiniuurl = url
						fmt.Printf("数据空间链接: %s\n", qiniuurl)

						fmt.Print("操作完成\n")
					}

				} else {
					fmt.Printf("%v\n", err)
				}
				if ok, err := satool.Createqiniubucket(child, uid,  QINIU_URM_PREFIX +qiniuname); ok {
					fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_URM_PREFIX  +qiniuname)
					url, err := satool.Getqiniubucket(child, QINIU_URM_PREFIX  +qiniuname)
					if err != nil {
						fmt.Printf("%v\n", err)
					} else {
						qiniuurl = url
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


func sharefamily() {
	//创建RDS账号
	//创建RDS库
	//创建RDS账号
	//创建RDS库
	//创建七牛账号
	var sitename string
	var rdsid, rdsusername, rdspassword, wuzhekadbname, urmdbname string
	var qiniuname, qiniupassword, qiniuemail, qiniuak, qiniusk, qiniuurl string
	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入RDS账号名:")
	fmt.Scanln(&rdsusername)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)
	wuzhekadbname = rdsusername + "-family"
	urmdbname = rdsusername + "-urm"

	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)
	qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX

	fmt.Print("正在操作RDS...\n")
	//创建rds账号
	if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername, rdspassword, sitename); ok {
		fmt.Printf("创建账号成功: %s[%s]\n", rdsusername, rdspassword)
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
		if accreated {
			//WUZHEKA
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname, sitename); ok {
				fmt.Printf("创建family数据库成功: %s\n", wuzhekadbname)
				dbcreated := false
				for !dbcreated {
					dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname)
					if err != nil {
						fmt.Printf("%v\n", err)
						break
					}
					switch dbstatus {
					case rds.DB_STATUS_RUNNING:
						dbcreated = true
					case rds.DB_STATUS_CREATING:
						fmt.Print("family数据库创建中，等待3s再次检查...\n")
						time.Sleep(3 * time.Second)
					case rds.DB_STATUS_DELETING:
						fmt.Print("family数据库删除中...\n")
						break
					}
				}
				if dbcreated {
					if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, wuzhekadbname, rdsusername); ok {
						fmt.Print("family授权成功\n")

					} else {
						fmt.Printf("%v\n", err)
					}
				}
			} else {
				fmt.Printf("%v\n", err)
			}

			//URM
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname, sitename); ok {
				fmt.Printf("创建URM数据库成功: %s\n", urmdbname)
				dbcreated := false
				for !dbcreated {
					dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname)
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
				if dbcreated {
					if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, urmdbname, rdsusername); ok {
						fmt.Print("URM授权成功\n")

					} else {
						fmt.Printf("%v\n", err)
					}
				}
			} else {
				fmt.Printf("%v\n", err)
			}
		}

	} else {
		fmt.Printf("%v\n", err)
	}

	if len(qiniuname) > 0 {
		fmt.Print("正在操作七牛...\n")
		access, err := satool.Getqiniuaccess(QINIU_USERNAME, QINIU_PASSWORD)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("获取七牛授权成功\n")
			qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX
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
						qiniuak = info.Key
						qiniusk = info.Secret
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
				if ok, err := satool.Createqiniubucket(child, uid, QINIU_FAMILY_PREFIX +qiniuname); ok {
					fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_FAMILY_PREFIX +qiniuname)
					url, err := satool.Getqiniubucket(child, QINIU_FAMILY_PREFIX +qiniuname)
					if err != nil {
						fmt.Printf("%v\n", err)
					} else {
						qiniuurl = url
						fmt.Printf("数据空间链接: %s\n", qiniuurl)

						fmt.Print("操作完成\n")
					}

				} else {
					fmt.Printf("%v\n", err)
				}
				if ok, err := satool.Createqiniubucket(child, uid,  QINIU_URM_PREFIX +qiniuname); ok {
					fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_URM_PREFIX  +qiniuname)
					url, err := satool.Getqiniubucket(child, QINIU_URM_PREFIX  +qiniuname)
					if err != nil {
						fmt.Printf("%v\n", err)
					} else {
						qiniuurl = url
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


//补充五折卡站点
func addwuzheka() {
	//创建RDS账号
	//创建RDS数据库
	//创建bucket
	var sitename string
	var rdsid, rdsdbname, rdspassword string
	var qiniuname string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入七牛名(不含前后缀):")
	fmt.Scanln(&qiniuname)

	rdsdbname = "hangjia_wuzheka"
	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)

	addsite("WUZHEKA", sitename, rdsid, rdsdbname, rdspassword, QINIU_WUZHEKA_PREFIX, qiniuname)
}

//舆情
func yuqing() {
	var sitename string
	var rdsid, rdsdbname, rdsusername, rdspassword string
	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入RDS账号名:")
	fmt.Scanln(&rdsusername)

	rdsdbname = "yuqing-" + rdsusername

	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Print("正在操作...\n")
	if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsdbname, sitename); ok {
		fmt.Printf("创建数据库成功: %s\n", rdsdbname)
		//创建rds账号
		if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername, rdspassword, sitename); ok {
			fmt.Printf("创建账号成功: %s[%s]\n", rdsusername, rdspassword)
			//检查数据库状态
			dbcreated := false
			for !dbcreated {
				dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsdbname)
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
				//授权
				if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsdbname, rdsusername); ok {
					fmt.Print("授权成功\n")

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
}
