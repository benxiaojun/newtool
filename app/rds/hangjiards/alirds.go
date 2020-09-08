package main

import (
	"fmt"
	"github.com/benxiaojun/satool/app/rds"
	"os"
)

const VERSION = "v1.0 20150725"

const (
	ALIYUN_ACCESSKEY_ID     = "YZ2RD2dX6DxJTMm9"
	ALIYUN_ACCESSKEY_SECRET = "T3PPRJ3WzAXBEsL2nBSzhMHkGyilQX"
)

func main() {
	fmt.Print("列出指定时间的慢查询或错误查询 - For 航加\n")

	var act string
	for {
		fmt.Scanf("%s\n", &act)
		if len(act) > 0 {
			switch act {
			case "q", "exit", "quit":
				fmt.Println("See you...")
				os.Exit(0)
			case "v", "version":
				fmt.Println(VERSION)
			case "h", "help":
				help()
			case "i", "instance":
				rds.Listinstances(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET)
			case "s", "slow":
				rds.Listslowlog(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET)
			case "e", "error":
				rds.Listerrorlog(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET)
			case "es", "exports":
				rds.Export(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, "slow")
			case "ee", "exporte":
				rds.Export(ALIYUN_ACCESSKEY_ID, ALIYUN_ACCESSKEY_SECRET, "error")
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
    
    i|instance                    显示RDS实例
    s|slow                        显示慢查询
    e|error                       显示错误查询
    es|exports                    导出慢查询日志
    ee|exporte                    导出错误日志
`)
}
