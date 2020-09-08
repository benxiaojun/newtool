package main

import (
	"fmt"
	"os"

	"github.com/benxiaojun/satool/app/qiniu/model"
	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gotime"
)

func main() {
	workdir, _ := gofile.WorkDir()
	model.Init(workdir, "2789378367@qq.com", "19850314@qq.com", 1)
	var act string
	fmt.Println("Welcome To Qianfan Qiniu Usage Tool")

	for {
		fmt.Scanf("%s\n", &act)
		if len(act) > 0 {
			switch act {
			case "q", "exit", "quit":
				fmt.Println("See you...")
				os.Exit(0)
			case "fetch":
				var month string
				fmt.Printf("输入月份[如: %s]:", gotime.GetTimeStr(gotime.Y+gotime.M))
				fmt.Scanf("%s\n", &month)
				if len(month) == 6 {
					model.Fetch(month)
				}
			case "semail":
				var email string
				fmt.Print("输入邮箱:")
				fmt.Scanf("%s\n", &email)
				if len(email) > 0 {
					model.Search(email, "")
				}
			case "smonth":
				var month string
				fmt.Printf("输入月份[如: %s]:", gotime.GetTimeStr(gotime.Y+gotime.M))
				fmt.Scanf("%s\n", &month)
				if len(month) == 6 {
					model.Search("", month)
				}
			case "se":
				model.Summaryemail()
			case "sm":
				model.Summarymonth()
			case "clean":
				var month string
				fmt.Printf("输入月份[如: %s]:", gotime.GetTimeStr(gotime.Y+gotime.M))
				fmt.Scanf("%s\n", &month)
				if len(month) == 6 {
					model.Clean(month)
				}
			case "h", "help":
				model.Usage()
			}
		}
	}
}
