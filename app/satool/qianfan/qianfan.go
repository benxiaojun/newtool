package main

import (
	"fmt"
	"github.com/benxiaojun/satool/aliyun/rds"
	"github.com/benxiaojun/satool/app/satool"
	"github.com/sillydong/goczd/godata"
	"github.com/sillydong/goczd/gohttp"
	"os"
	"time"
)

const DEBUG = false
const VERSION = "v1.0 20150725"

const (
	DNSPOD_TOKEN            = ""
	DNSPOD_QIANFANAPI       = ""
	DNSPOD_QIANFANYUN       = ""
	QINIU_USERNAME          = ""
	QINIU_PASSWORD          = ""
	QINIU_EMAIL_PREFIX      = "client_"
	QINIU_EMAIL_SUFFIX      = "@qianfanyun.com"
	QINIU_BUCKET_PREFIX     = "qianfanyun-"
	ALIYUN_ACCESSKEY_ID     = ""
	ALIYUN_ACCESSKEY_SECRET = ""
)

func main() {
	var act string
	fmt.Printf("Welcome to SATOOL %s For QianFan\nType h to view available commands\n", VERSION)

	gohttp.REQUEST_DEBUG = DEBUG

	for {
		fmt.Scanln(&act)
		if len(act) > 0 {
			switch act {
			case "q", "exit", "quit":
				fmt.Println("See you...")
				os.Exit(0)
			case "v", "version":
				fmt.Println(VERSION)
			case "h", "help":
				help()
			case "zunxiang":
				zunxiang()
			case "sharexinrui":
				sharexinrui()
			case "monoxinrui":
				monoxinrui()
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
    v|version                     打印程序版本
    h|help                        打印帮助信息
    q|exit|quit                   退出程序
    
    zunxiang                      尊享
    sharexinrui                   共享新锐
    monoxinrui                    独享新锐
    t                             计算日流量
`)
}

//创建尊享
func zunxiang() {
	var sitename string
	var region string
	var ecsid, ecspublicip, ecsinternalip string
	var rdsid, rdsurl, rdsusername, rdsdbname, rdspassword string
	var qiniuname, qiniupassword, qiniuemail, qiniuak, qiniusk, qiniuurl string
	var dnsprefix string

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

	fmt.Print("输入RDS账号名:")
	fmt.Scanln(&rdsusername)

	fmt.Print("输入RDS库名:")
	fmt.Scanln(&rdsdbname)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	fmt.Print("输入DNS域名前缀:")
	fmt.Scanln(&dnsprefix)

	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)
	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Println("正在操作...")

	//更新ECS备注
	satool.Modifyecs(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, ecsid, sitename)
	//获取ecs信息
	ecs, err := satool.Getecs(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, ecsid, region)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		ecspublicip = ecs.PublicIpAddress.IpAddress[0]
		ecsinternalip = ecs.InnerIpAddress.IpAddress[0]
		fmt.Printf("ECS外网IP: %s\nECS内网IP: %s\n", ecspublicip, ecsinternalip)

		//更新RDS备注
		satool.Modifyrds(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, sitename)
		//获取rds信息
		rdsinfo, err := satool.Getrds(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			rdsurl = rdsinfo.ConnectionString
			fmt.Printf("RDS地址: %s\n", rdsurl)
			//创建rds数据库
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsdbname, sitename); ok {
				fmt.Printf("创建RDS数据库成功: %s\n", rdsdbname)
				//创建rds账号
				if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername, rdspassword, sitename); ok {
					fmt.Printf("创建RDS账号成功: %s[%s]\n", rdsusername, rdspassword)
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
							//设置白名单
							if ok, err := satool.Modifysecurityips(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, ecsinternalip); ok {
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
											qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX
											fmt.Printf("七牛子账号: %s[%s]\n", qiniuemail, qiniupassword)
											//创建七牛子账号
											qiniu, err := satool.Createqiniuaccount(*access, QINIU_EMAIL_PREFIX+qiniuname+QINIU_EMAIL_SUFFIX, qiniupassword)
											if err != nil {
												fmt.Printf("%v\n", err)
											} else {
												fmt.Printf("子账号创建成功: %d\n", qiniu.UID)
												info, err := satool.Getqiniukey(*access, qiniu.UID)
												if err != nil {
													fmt.Printf("%v\n", err)
												} else {
													qiniuak = info.Key
													qiniusk = info.Secret
													fmt.Printf("AK: %s\nSK: %s\n", qiniuak, qiniusk)
													child, err := satool.Getqiniuchild(qiniu.UID, *access)
													if err != nil {
														fmt.Printf("%v\n", err)
													} else {
														fmt.Print("获取子账号授权成功\n")
														//创建bucket
														if ok, err := satool.Createqiniubucket(child, qiniu.UID, QINIU_BUCKET_PREFIX+qiniuname); ok {
															fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_BUCKET_PREFIX+qiniuname)
															url, err := satool.Getqiniubucket(child, QINIU_BUCKET_PREFIX+qiniuname)
															if err != nil {
																fmt.Printf("%v\n", err)
															} else {
																qiniuurl = url
																fmt.Printf("数据空间链接: %s\n", qiniuurl)
															}

														} else {
															fmt.Printf("%v\n", err)
														}
													}
												}
											}
										}
									} else {
										fmt.Println("跳过七牛")
									}

									//创建dnspod
									if len(dnsprefix) > 0 {
										//x.qianfanapi.com
										if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, dnsprefix, ecspublicip); ok {
											fmt.Printf("域名解析成功[%s]%s\n", dnsprefix, ecspublicip)
										} else {
											fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
										}

										//x.pic.qianfanapi.com
										if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, dnsprefix+".pic", ecspublicip); ok {
											fmt.Printf("域名解析[%s]%s\n", dnsprefix+".pic", ecspublicip)
										} else {
											fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
										}

										//x.adm.qianfanyun.com
										if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, dnsprefix+".adm", ecspublicip); ok {
											fmt.Printf("域名解析[%s]%s\n", dnsprefix+".adm", ecspublicip)
										} else {
											fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
										}

										//x.share.qianfanyun.com
										if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, dnsprefix+".share", ecspublicip); ok {
											fmt.Printf("域名解析[%s]%s\n", dnsprefix+".share", ecspublicip)
										} else {
											fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
										}
									} else {
										fmt.Println("跳过DNS解析")
									}

									fmt.Print("操作完成\n")
								}
							} else {
								fmt.Printf("%v\n", err)
							}
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
	}
}

func sharexinrui() {
	var sitename string
	var region string
	var ecsid, ecspublicip, ecsinternalip string
	var rdsid, rdsurl, rdsusername, rdsdbname, rdspassword string
	var qiniuname, qiniupassword, qiniuemail, qiniuak, qiniusk, qiniuurl string
	var dnsprefix string

	fmt.Print("输入站点名:")
	fmt.Scanln(&sitename)

	fmt.Print("输入ECS ID:")
	fmt.Scanln(&ecsid)

	fmt.Print("输入ECS地域[cn-hangzhou]:")
	fmt.Scanln(&region)
	if len(region) == 0 {
		region = "cn-hangzhou"
	}

	fmt.Print("输入RDS ID:")
	fmt.Scanln(&rdsid)

	fmt.Print("输入RDS账号名:")
	fmt.Scanln(&rdsusername)

	fmt.Print("输入RDS库名:")
	fmt.Scanln(&rdsdbname)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	fmt.Print("输入DNS域名前缀:")
	fmt.Scanln(&dnsprefix)

	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)
	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Println("正在操作...")
        if len(qiniuname) > 0 {
                access, err := satool.Getqiniuaccess(QINIU_USERNAME, QINIU_PASSWORD)
                if err != nil {
                        fmt.Printf("%v\n", err)
                } else {
                        fmt.Print("获取七牛授权成功\n")
                        qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX
                        fmt.Printf("七牛子账号: %s[%s]\n", qiniuemail, qiniupassword)
                        //创建七牛子账号
                        qiniu, err := satool.Createqiniuaccount(*access, QINIU_EMAIL_PREFIX+qiniuname+QINIU_EMAIL_SUFFIX, qiniupassword)
                        if err != nil {
                                fmt.Printf("%v\n", err)
                        } else {
                                fmt.Printf("子账号创建成功: %d\n", qiniu.UID)
                                info, err := satool.Getqiniukey(*access, qiniu.UID)
                                if err != nil {
                                        fmt.Printf("%v\n", err)
                                } else {
                                        qiniuak = info.Key
                                        qiniusk = info.Secret
                                        fmt.Printf("AK: %s\nSK: %s\n", qiniuak, qiniusk)
                                        child, err := satool.Getqiniuchild(qiniu.UID, *access)
                                        if err != nil {
                                                fmt.Printf("%v\n", err)
                                        } else {
                                                fmt.Print("获取子账号授权成功\n")
                                                //创建bucket
                                                if ok, err := satool.Createqiniubucket(child, qiniu.UID, QINIU_BUCKET_PREFIX+qiniuname); ok {
                                                        fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_BUCKET_PREFIX+qiniuname)
                                                        url, err := satool.Getqiniubucket(child, QINIU_BUCKET_PREFIX+qiniuname)
                                                        if err != nil {
                                                                fmt.Printf("%v\n", err)
                                                        } else {
                                                                qiniuurl = url
                                                                fmt.Printf("数据空间链接: %s\n", qiniuurl)
                                                        }

                                                } else {
                                                        fmt.Printf("%v\n", err)
                                                }
                                        }
                                }
                        }
                }
        } else {
                fmt.Println("跳过七牛")
        }

        //创建dnspod
        if len(dnsprefix) > 0 {
                //x.qianfanapi.com
                if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, dnsprefix, ecspublicip); ok {
                        fmt.Printf("域名解析成功[%s]%s\n", dnsprefix, ecspublicip)
                } else {
                        fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
                }

                //x.pic.qianfanapi.com
                if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, dnsprefix+".pic", ecspublicip); ok {
                        fmt.Printf("域名解析[%s]%s\n", dnsprefix+".pic", ecspublicip)
                } else {
                        fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
                }

                //x.adm.qianfanyun.com
                if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, dnsprefix+".adm", ecspublicip); ok {
                        fmt.Printf("域名解析[%s]%s\n", dnsprefix+".adm", ecspublicip)
                } else {
                        fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
                }

                //x.share.qianfanyun.com
                if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, dnsprefix+".share", ecspublicip); ok {
                        fmt.Printf("域名解析[%s]%s\n", dnsprefix+".share", ecspublicip)
                } else {
                        fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
                }
        } else {
                fmt.Println("跳过DNS解析")
        }

        fmt.Print("操作完成\n")

}

func monoxinrui() {
	var sitename string
	var region string
	var ecsid, ecspublicip, ecsinternalip string
	var rdsid, rdsurl, rdsusername, rdsdbname, rdspassword string
	var qiniuname, qiniupassword, qiniuemail, qiniuak, qiniusk, qiniuurl string
	var dnsprefix string

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

	fmt.Print("输入RDS账号名:")
	fmt.Scanln(&rdsusername)

	fmt.Print("输入RDS库名:")
	fmt.Scanln(&rdsdbname)

	fmt.Print("输入七牛子账号名:")
	fmt.Scanln(&qiniuname)

	fmt.Print("输入DNS域名前缀:")
	fmt.Scanln(&dnsprefix)

	rdspassword = godata.RandomString(12, godata.ALPHANUMERIC)
	qiniupassword = godata.RandomString(12, godata.ALPHANUMERIC)

	fmt.Print("正在操作...\n")

	//更新ECS备注
	satool.Modifyecs(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, ecsid, sitename)
	//获取ecs信息
	ecs, err := satool.Getecs(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, ecsid, region)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		ecspublicip = ecs.PublicIpAddress.IpAddress[0]
		ecsinternalip = ecs.InnerIpAddress.IpAddress[0]
		fmt.Printf("ECS外网IP: %s\nECS内网IP: %s\n", ecspublicip, ecsinternalip)

		//更新RDS备注
		satool.Modifyrds(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, sitename)
		//获取rds信息
		rdsinfo, err := satool.Getrds(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			rdsurl = rdsinfo.ConnectionString
			fmt.Printf("RDS地址: %s\n", rdsurl)
			//创建千帆rds数据库
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsdbname, sitename); ok {
				fmt.Printf("创建千帆RDS数据库成功: %s\n", rdsdbname)
				//创建千帆rds账号
				if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsusername, rdspassword, sitename); ok {
					fmt.Printf("创建千帆RDS账号成功: %s[%s]\n", rdsusername, rdspassword)
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
				} else {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Printf("%v\n", err)
			}

			//创建bbs rds数据库
			if ok, err := satool.Createrdsdatabase(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, "bbs_"+rdsdbname, sitename+"论坛"); ok {
				fmt.Printf("创建论坛RDS数据库成功: %s\n", "bbs_"+rdsdbname)
				//创建bbs rds账号
				if ok, err := satool.Createrdsaccount(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, "bbs_"+rdsusername, rdspassword, sitename+"论坛"); ok {
					fmt.Printf("创建论坛RDS账号成功: %s[%s]\n", "bbs_"+rdsusername, rdspassword)
					//检查数据库状态
					dbcreated := false
					for !dbcreated {
						dbstatus, err := satool.Getdatabasestatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, "bbs_"+rdsdbname)
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
						acstatus, err := satool.Getaccountstatus(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, "bbs_"+rdsusername)
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
				} else {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Printf("%v\n", err)
			}

			//授权
			if ok, err := satool.Grantrdsprivilege(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, rdsdbname, rdsusername); ok {
				fmt.Print("授权成功\n")
				//设置白名单
				if ok, err := satool.Modifysecurityips(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, rdsid, ecsinternalip); ok {
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
								qiniuemail = QINIU_EMAIL_PREFIX + qiniuname + QINIU_EMAIL_SUFFIX
								fmt.Printf("七牛子账号: %s[%s]\n", qiniuemail, qiniupassword)
								//创建七牛子账号
								qiniu, err := satool.Createqiniuaccount(*access, QINIU_EMAIL_PREFIX+qiniuname+QINIU_EMAIL_SUFFIX, qiniupassword)
								if err != nil {
									fmt.Printf("%v\n", err)
								} else {
									fmt.Printf("子账号创建成功: %d\n", qiniu.UID)
									info, err := satool.Getqiniukey(*access, qiniu.UID)
									if err != nil {
										fmt.Printf("%v\n", err)
									} else {
										qiniuak = info.Key
										qiniusk = info.Secret
										fmt.Printf("AK: %s\nSK: %s\n", qiniuak, qiniusk)
										child, err := satool.Getqiniuchild(qiniu.UID, *access)
										if err != nil {
											fmt.Printf("%v\n", err)
										} else {
											fmt.Print("获取子账号授权成功\n")
											//创建bucket
											if ok, err := satool.Createqiniubucket(child, qiniu.UID, QINIU_BUCKET_PREFIX+qiniuname); ok {
												fmt.Printf("创建子账号数据空间[%s]成功\n", QINIU_BUCKET_PREFIX+qiniuname)
												url, err := satool.Getqiniubucket(child, QINIU_BUCKET_PREFIX+qiniuname)
												if err != nil {
													fmt.Printf("%v\n", err)
												} else {
													qiniuurl = url
													fmt.Printf("数据空间链接: %s\n", qiniuurl)
												}
											} else {
												fmt.Printf("%v\n", err)
											}
										}
									}
								}
							}
						} else {
							fmt.Println("跳过七牛")
						}

						//创建dnspod
						if len(dnsprefix) > 0 {
							//x.qianfanapi.com
							if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, dnsprefix, ecspublicip); ok {
								fmt.Printf("域名解析成功[%s]%s\n", dnsprefix, ecspublicip)
							} else {
								fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
							}

							//x.pic.qianfanapi.com
							if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, dnsprefix+".pic", ecspublicip); ok {
								fmt.Printf("域名解析[%s]%s\n", dnsprefix+".pic", ecspublicip)
							} else {
								fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
							}

							//x.adm.qianfanyun.com
							if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, dnsprefix+".adm", ecspublicip); ok {
								fmt.Printf("域名解析[%s]%s\n", dnsprefix+".adm", ecspublicip)
							} else {
								fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
							}

							//x.share.qianfanyun.com
							if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, dnsprefix+".share", ecspublicip); ok {
								fmt.Printf("域名解析[%s]%s\n", dnsprefix+".share", ecspublicip)
							} else {
								fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
							}

							//x.bbs.qianfanyun.com
							if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, dnsprefix+".bbs", ecspublicip); ok {
								fmt.Printf("域名解析[%s]%s\n", dnsprefix+".share", ecspublicip)
							} else {
								fmt.Printf("域名解析失败[%s]%s: %v\n", dnsprefix, ecspublicip, err)
							}
						} else {
							fmt.Println("跳过DNS解析")
						}

						fmt.Print("操作完成\n")
					}
				} else {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Printf("%v\n", err)
			}
		}
	}
}
