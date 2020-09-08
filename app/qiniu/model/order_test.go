package model

import (
	"fmt"
	"github.com/sillydong/goczd/gotime"
	"testing"
	"time"
)

func init() {
	fmt.Println("init db")
	if err := LoadConfig("/Volumes/DATA/Src/Go/src/github.com/benxiaojun/satool/app/qiniu/db.json"); err != nil {
		fmt.Printf("%v", err)
	} else if err := InitEngine(); err != nil {
		fmt.Printf("%v", err)
	}
}

/*
func Test_GetOrders(t *testing.T){
	orders,err:= GetOrders("client_18qiang@qianfanyun.com","")
	if err != nil {
		t.Error(err)
	}else{
		for index,order := range orders {
			fmt.Printf("%d %+v\n",index,order)
		}
	}
}
*/

/*
func Test_GetOrders2(t *testing.T){
	orders,err:=GetOrders("","201505")
	if err != nil {
		t.Error(err)
	}else {
		for index, order := range orders {
			fmt.Printf("%d %+v\n", index, order)
		}
	}
}
*/

/*
func Test_GetOrders3(t *testing.T) {
	orders, err := GetOrders("client_18qiang@qianfanyun.com", "201505")
	if err != nil {
		t.Error(err)
	}else {
		for index, order := range orders {
			fmt.Printf("%d %+v\n", index, order)
		}
	}
}
*/

/*
func Test_SumOrders(t *testing.T){
	info,err:=SumOrders("client_18qiang@qianfanyun.com","")
	if err != nil {
		t.Error(err)
	}else{
		for _,item:=range info{
			for key,val :=range item{
				fmt.Printf("%s: %s\n", key, string(val))
			}
		}
	}
}
*/

/*
func Test_SumOrders2(t *testing.T) {
	info, err := SumOrders("", "201505")
	if err != nil {
		t.Error(err)
	}else {
		fmt.Printf("%+v\n", info)
	}
}

func Test_SumOrders3(t *testing.T){
	info, err := SumOrders("client_18qiang@qianfanyun.com", "201505")
	if err != nil {
		t.Error(err)
	}else {
		fmt.Printf("%+v\n", info)
	}
}

*/

func TestDate(t *testing.T) {
	now, _ := time.Parse(gotime.FORMAT_YYYY_MM_DD, time.Now().Format(gotime.FORMAT_YYYY_MM)+"-01")
	fmt.Println(now)
	fmt.Println(now.AddDate(0, 1, 0))
	fmt.Println(now.AddDate(0, 2, 0))
	fmt.Println(now.AddDate(0, -1, 0))
	fmt.Println(now.AddDate(0, -2, 0))
}
