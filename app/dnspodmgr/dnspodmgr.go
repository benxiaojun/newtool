package main

import (
	"fmt"
	"github.com/benxiaojun/satool/app/satool"
	"os"
)

const VERSION = "v1.0 20160302"
const (
	DNSPOD_TOKEN      = "10417,1418b4b58448649383dc627382ede011"
	DNSPOD_QIANFANAPI = "22010540"
	DNSPOD_QIANFANYUN = "22010518"
)

func main() {
	var act string
	fmt.Printf("Welcome to DNSPODMGR %s For QianFan\nType h to view available commands\n", VERSION)

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
			case "zunxiang":
				zunxiang()
			case "monoxinrui":
				monoxinrui()
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
    
    zunxiang                      尊享(x, x.pic, x.adm, x.share)
    monoxinrui                    独享新锐(x, x.pic, x.adm, x.share, x.bbs)
`)
}

func zunxiang() {
	var prefix, ip, confirm string

	fmt.Print("请输入前缀:")
	fmt.Scanf("%s\n", &prefix)

	fmt.Print("请输入IP:")
	fmt.Scanf("%s\n", &ip)

	if len(prefix) > 0 && len(ip) > 0 {
		fmt.Println("确认创建以下域名?")
		fmt.Printf("  %s.qianfanapi.com\n  %s.pic.qianfanapi.com\n  %s.adm.qianfanyun.com\n  %s.share.qianfanyun.com\n", prefix, prefix, prefix, prefix)
		fmt.Printf("指向: %s\n", ip)

		fmt.Print("[y|n]:")
		fmt.Scanf("%s\n", &confirm)
		if confirm == "y" {
			//x.qianfanapi.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, prefix, ip); ok {
				fmt.Printf("域名解析成功[%s]%s\n", prefix, ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}

			//x.pic.qianfanapi.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, prefix+".pic", ip); ok {
				fmt.Printf("域名解析[%s]%s\n", prefix+".pic", ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}

			//x.adm.qianfanyun.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, prefix+".adm", ip); ok {
				fmt.Printf("域名解析[%s]%s\n", prefix+".adm", ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}

			//x.share.qianfanyun.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, prefix+".share", ip); ok {
				fmt.Printf("域名解析[%s]%s\n", prefix+".share", ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}
		}
	}
}

func monoxinrui() {
	var prefix, ip, confirm string

	fmt.Print("请输入前缀:")
	fmt.Scanf("%s\n", &prefix)

	fmt.Print("请输入IP:")
	fmt.Scanf("%s\n", &ip)

	if len(prefix) > 0 && len(ip) > 0 {
		fmt.Println("确认创建以下域名?")
		fmt.Printf("  %s.qianfanapi.com\n  %s.pic.qianfanapi.com\n  %s.adm.qianfanyun.com\n  %s.share.qianfanyun.com\n", prefix, prefix, prefix, prefix)
		fmt.Printf("指向: %s\n", ip)

		fmt.Print("[y|n]:")
		fmt.Scanf("%s\n", &confirm)
		if confirm == "y" {
			//x.qianfanapi.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, prefix, ip); ok {
				fmt.Printf("域名解析成功[%s]%s\n", prefix, ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}

			//x.pic.qianfanapi.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANAPI, prefix+".pic", ip); ok {
				fmt.Printf("域名解析[%s]%s\n", prefix+".pic", ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}

			//x.adm.qianfanyun.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, prefix+".adm", ip); ok {
				fmt.Printf("域名解析[%s]%s\n", prefix+".adm", ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}

			//x.share.qianfanyun.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, prefix+".share", ip); ok {
				fmt.Printf("域名解析[%s]%s\n", prefix+".share", ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}

			//x.share.qianfanyun.com
			if ok, err := satool.Creatednspod(DNSPOD_TOKEN, DNSPOD_QIANFANYUN, prefix+".bbs", ip); ok {
				fmt.Printf("域名解析[%s]%s\n", prefix+".bbs", ip)
			} else {
				fmt.Printf("域名解析失败[%s]%s: %v\n", prefix, ip, err)
			}
		}
	}
}
