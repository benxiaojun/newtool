package model

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/benxiaojun/satool/newqiniu"

	"github.com/olebedev/config"
	simplejson "github.com/sillydong/go-simplejson"
	"github.com/sillydong/goczd/godata"
	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gohttp"
	"github.com/sillydong/goczd/gotime"
)

var workdir string
var username string = "2789378367@qq.com"
var password string = "19850314@qq.com"
var company int = 1

func Init(iworkdir, iusername, ipassword string, icompany int) {
	workdir = iworkdir
	username = iusername
	password = ipassword
	company = icompany
}

func Usage() {
	fmt.Println(`操作列表:
    h|help                        打印帮助信息
    q|exit|quit                   退出程序
    
    fetch                         拉取指定月份子账号用量
    semail                        统计指定邮箱按月份用量
    smonth                        统计指定月份各子账号用量
    se                            所有子账号费用统计信息
    sm                            所有月份费用统计信息
    clean                         清理指定月份数据
`)
}

func Fetch(month string) {
	if err := LoadConfig(workdir + "/db.json"); err != nil {
		fmt.Printf("%v", err)
		return
	} else if err := InitEngine(); err != nil {
		fmt.Printf("%v", err)
		return
	}

	priceconf, err := getpriceconf()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	current, _ := strconv.Atoi(gotime.GetTimeStr(gotime.Y + gotime.M))
	request, _ := strconv.Atoi(month)
	save := true
	if current > request {
		save = true
	}

	access, err := getqiniuaccess(username, password)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {

		fmt.Print("获取主账号用量...\n")
		usage, err := getaccountusage(access, month)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			if save {
				saveaccountorder(access, month, usage, priceconf)
			} else {
				fmt.Printf("%s: Space[%v] Get[%v] Put[%v] Transfer[%v]\n", access.Email, godata.FriendlyByte(usage.Space), 0, usage.Put, godata.FriendlyByte(usage.Traffic))
			}
		}

		fmt.Print("获取子账号用量...\n")
		children, err := getchildren(access)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			for _, child := range children {
				usage, err := getchildusage(child, month)
				if err != nil {
					fmt.Printf("%v\n", err)
				} else {
					if save {
						savechildorder(child, month, usage, priceconf)
					} else {
						fmt.Printf("%s: Space[%v] Get[%v] Put[%v] Transfer[%v]\n", child.Email, godata.FriendlyByte(usage.Space), 0, usage.Put, godata.FriendlyByte(usage.Traffic))
					}
				}
			}
		}
	}
}

func Clean(month string) {
	if err := LoadConfig(workdir + "/db.json"); err != nil {
		fmt.Printf("%v", err)
		return
	}
	if err := InitEngine(); err != nil {
		fmt.Printf("%v", err)
		return
	}

	num, err := DropMonth(company, month)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("删除%d条记录\n", num)
	}
}

func Search(email, month string) {
	if err := LoadConfig(workdir + "/db.json"); err != nil {
		fmt.Printf("%v", err)
		return
	}
	if err := InitEngine(); err != nil {
		fmt.Printf("%v", err)
		return
	}

	summary, err := SumOrders(company, email, month)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("总计: Space[%v] Get[%v] Put[%v] Transfer[%v] Price[%v]\n", string(summary[0]["sum_space"]), string(summary[0]["sum_apicall_get"]), string(summary[0]["sum_apicall_put"]), string(summary[0]["sum_transfer"]), string(summary[0]["sum_price"]))
	}

	fmt.Println("明细:")
	orders, err := GetOrders(company, email, month)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, order := range orders {
			fmt.Printf("[%d]%s[%s]: Space[%v] Get[%v] Put[%v] Transfer[%v] Price[%v]\n", order.Id, order.Email, order.Month, godata.FriendlyByte(order.Space), order.ApicallGet, order.ApicallPut, godata.FriendlyByte(order.Transfer), order.Price)
		}
	}
}

func Summaryemail() {
	if err := LoadConfig(workdir + "/db.json"); err != nil {
		fmt.Printf("%v", err)
		return
	}
	if err := InitEngine(); err != nil {
		fmt.Printf("%v", err)
		return
	}

	access, err := getqiniuaccess(username, password)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		children, err := getchildren(access)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			for _, child := range children {
				summary, err := SumOrders(company, child.Email, "")
				if err != nil {
					fmt.Printf("%v\n", err)
				} else {
					friendlyspace, _ := strconv.ParseInt(string(summary[0]["sum_space"]), 10, 64)
					friendlytransfer, _ := strconv.ParseInt(string(summary[0]["sum_transfer"]), 10, 64)
					fmt.Printf("%s: Space[%v] Get[%v] Put[%v] Transfer[%v] Price[%v]\n", child.Email, godata.FriendlyByte(friendlyspace), string(summary[0]["sum_apicall_get"]), string(summary[0]["sum_apicall_put"]), godata.FriendlyByte(friendlytransfer), string(summary[0]["sum_price"]))
					//fmt.Printf("%s\t%v\t%v\t%v\t%v\t%v\n", child.Email, godata.FriendlyByte(friendlyspace), string(summary[0]["sum_apicall_get"]), string(summary[0]["sum_apicall_put"]), godata.FriendlyByte(friendlytransfer), string(summary[0]["sum_price"]))
				}
			}
		}
	}
}

func Summarymonth() {
	if err := LoadConfig(workdir + "/db.json"); err != nil {
		fmt.Printf("%v", err)
		return
	}
	if err := InitEngine(); err != nil {
		fmt.Printf("%v", err)
		return
	}

	start, _ := time.Parse(gotime.Y+gotime.M+gotime.D, "20150301")
	now, _ := time.Parse(gotime.FORMAT_YYYY_MM_DD, time.Now().Format(gotime.FORMAT_YYYY_MM)+"-01")
	diff := 0

	for {
		month := now.AddDate(0, diff, 0)
		if month.Before(start) {
			break
		}

		summary, err := SumOrders(company, "", month.Format(gotime.Y+gotime.M))
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			friendlyspace, _ := strconv.ParseInt(string(summary[0]["sum_space"]), 10, 64)
			friendlytransfer, _ := strconv.ParseInt(string(summary[0]["sum_transfer"]), 10, 64)
			fmt.Printf("%s: Space[%v] Get[%v] Put[%v] Transfer[%v] Price[%v]\n", month.Format(gotime.Y+gotime.M), godata.FriendlyByte(friendlyspace), string(summary[0]["sum_apicall_get"]), string(summary[0]["sum_apicall_put"]), godata.FriendlyByte(friendlytransfer), string(summary[0]["sum_price"]))
		}

		diff--
	}
}

func getpriceconf() (*config.Config, error) {
	file := workdir + "/price.json"
	if gofile.FileExists(file) {
		return config.ParseJsonFile(file)
	}
	return nil, fmt.Errorf("file not exists")
}

type UsageData struct {
	Space   int64
	Put     int64
	Traffic int64
}

func getqiniuaccess(username, password string) (*newqiniu.Access, error) {
	access := &newqiniu.Access{}
	access.Username = username
	access.Password = password
	if gofile.FileExists(newqiniu.GetDir() + string(os.PathSeparator) + access.GetFileName()) {
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

func getchildren(access *newqiniu.Access) ([]*newqiniu.Child, error) {
	var children []*newqiniu.Child
	req := &newqiniu.ListChildAccountRequest{
		Offset: 0,
		Limit:  1000,
	}
	req.AuthorizationValue = access.AccessToken
	bytes, err := gohttp.DoGet(req)
	if err != nil {
		return nil, err
	} else {
		if string(bytes[0]) == "[" {
			//array
			resp := []newqiniu.ListChildAccountResponse{}
			err := json.Unmarshal(bytes, &resp)
			if err != nil {
				return nil, err
			} else {
				for _, cr := range resp {
					if !cr.IsDisabled {
						child := &newqiniu.Child{
							Uid:    cr.Uid,
							Email:  cr.Email,
							Userid: cr.Userid,
						}
						if gofile.FileExists(newqiniu.GetDir() + string(os.PathSeparator) + child.GetFileName()) {
							err := child.ReadKeys(true, access.AccessToken)
							if err != nil {
								return nil, err
							} else {
								children = append(children, child)
							}
						} else {
							err := child.Get(access.AccessToken)
							if err != nil {
								return nil, err
							} else {
								children = append(children, child)
							}
						}
					}
				}
				return children, nil
			}
		} else {
			//object
			resp := &newqiniu.ListChildAccountResponse{}
			err := json.Unmarshal(bytes, resp)
			if err != nil {
				return nil, err
			} else {
				return nil, fmt.Errorf("%v", resp.Msg())
			}
		}
	}
}

func getaccountusage(access *newqiniu.Access, month string) (*UsageData, error) {
	return getusage(access.Key, access.Secret, month)
}

func getchildusage(child *newqiniu.Child, month string) (*UsageData, error) {
	return getusage(child.Key, child.Secret, month)
}

func getusage(key, secret, month string) (*UsageData, error) {
	time.Sleep(time.Duration(1) * time.Second)
	usage := &UsageData{}

	_, ms, me, _ := gotime.CalcMonth(month)

	//put
	preq := &newqiniu.UsagePutRequest{
		Begin:       ms + "01/00:00",
		End:         me + "01/00:00",
		Granularity: newqiniu.GRANULARITY_DAY,
	}
	preq.AccessKey = key
	preq.SecretKey = secret
	pbytes, err := gohttp.DoGet(preq)
	if err != nil {
		return nil, err
	} else {
		pjson, err := simplejson.NewJson(pbytes)
		if err != nil {
			return nil, err
		} else {
			if pjson.IsObject() {
				return nil, fmt.Errorf("%v", pjson.Get("error").MustString())
			} else {
				for _, item := range pjson.MustArray() {
					hits, _ := (item.(map[string]interface{})["values"]).(map[string]interface{})["hits"].(json.Number).Int64()
					usage.Put += hits
				}
			}
		}
	}

	//space
	sreq := &newqiniu.UsageSpaceRequest{
		Begin:       ms + "28000000",
		End:         me + "01000000",
		Granularity: newqiniu.GRANULARITY_DAY,
	}
	sreq.AccessKey = key
	sreq.SecretKey = secret
	sresp := &newqiniu.UsageSpaceResponse{}
	err2 := gohttp.DoGetResponse(sreq, sresp)
	if err2 != nil {
		return nil, err2
	} else {
		if sresp.OK() {
			for _, num := range sresp.Datas {
				usage.Space = int64(math.Max(float64(usage.Space), float64(num)))
			}
		} else {
			return nil, fmt.Errorf("%v", sresp.Msg())
		}
	}

	//buckets
	breq := &newqiniu.ListBucketRequest{}
	breq.AccessKey = key
	breq.SecretKey = secret

	bytes, err := gohttp.DoPost(breq)
	if err != nil {
		return nil, err
	} else {
		if string(bytes[0]) == "[" {
			var bresp []string
			err := json.Unmarshal(bytes, &bresp)
			if err != nil {
				return nil, err
			} else {
				var domains []string

				for _, bucket := range bresp {
					//domains
					dreq := &newqiniu.ListBucketDomainRequest{}
					dreq.BucketName = bucket
					dreq.AccessKey = key
					dreq.SecretKey = secret
					dbytes, err := gohttp.DoGet(dreq)
					if err != nil {
						return nil, err
					} else {
						if string(dbytes[0]) == "[" {
							var dresp []string
							err := json.Unmarshal(dbytes, &dresp)
							if err != nil {
								return nil, err
							} else {
								domains = append(domains, dresp...)
							}
						} else {
							return nil, fmt.Errorf("%v", string(dbytes))
						}
					}
				}
				//traffic
				start, stop := monthstartstop(month, "%v-%02d-%v")
				if len(domains) > 0 {
					treq := &newqiniu.DomainTrafficRequest{
						Domains:     strings.Join(domains, ";"),
						Granularity: newqiniu.GRANULARITY_DAY,
						StartDate:   start,
						EndDate:     stop,
					}
					treq.AccessKey = key
					treq.SecretKey = secret
					tresp := &newqiniu.DomainTrafficResponse{}
					err := gohttp.DoPostJsonResponse(treq, tresp)
					if err != nil {
						return nil, err
					} else {
						if tresp.OK() {
							datas := tresp.Data
							for _, data := range datas {
								for _, traffics := range data {
									for _, traffic := range traffics {
										usage.Traffic += traffic
									}
								}
							}
						} else {
							return nil, fmt.Errorf("%v", tresp.Msg())
						}
					}
				}
			}
		} else {
			return nil, fmt.Errorf("%v", string(bytes))
		}
	}
	return usage, nil
}

func saveaccountorder(access *newqiniu.Access, month string, usage *UsageData, priceconf *config.Config) {
	saveorder(1, access.Uid, access.Userid, access.Email, month, usage, priceconf)
}

func savechildorder(child *newqiniu.Child, month string, usage *UsageData, priceconf *config.Config) {
	saveorder(2, child.Uid, child.Userid, child.Email, month, usage, priceconf)
}

func saveorder(t int, uid int64, userid string, email string, month string, usage *UsageData, priceconf *config.Config) {
	order := &Order{}
	order.Company = company
	order.Type = t
	order.Uid = int(uid)
	order.Userid = userid
	order.Email = email
	order.Month = month

	order.Space = usage.Space
	order.SpaceAvg = 0
	order.ApicallGet = 0
	order.ApicallPut = usage.Put
	order.Transfer = usage.Traffic

	order.PriceSpaceUnit = priceconf.UInt("price_space_unit", 1073741824)
	order.PriceSpace = priceconf.UFloat64("price_space", 0.165)
	order.PriceApicallGetUnit = priceconf.UInt("price_apicall_get_unit", 1000)
	order.PriceApicallGet = priceconf.UFloat64("price_apicall_get", 0.001)
	order.PriceApicallPutUnit = priceconf.UInt("price_apicall_put_unit", 1000)
	order.PriceApicallPut = priceconf.UFloat64("price_apicall_put", 0.01)
	order.PriceTransferUnit = priceconf.UInt("price_transfer_unit", 1073741824)
	order.PriceTransfer = priceconf.UFloat64("price_transfer", 0.25)

	price := ((float64(order.Space) / float64(order.PriceSpaceUnit) * order.PriceSpace) + (float64(order.ApicallGet) / float64(order.PriceApicallGetUnit) * order.PriceApicallGet) + (float64(order.ApicallPut) / float64(order.PriceApicallPutUnit) * order.PriceApicallPut) + (float64(order.Transfer) / float64(order.PriceTransferUnit) * order.PriceTransfer))
	if price > 0 {
		order.Price = strconv.FormatFloat(price, 'f', 4, 64)

		order.Importtime = int(gotime.GetTimestamp())
		order.Paytime = 0
		order.Invoicetime = 0

		err := CreateOrder(order)
		if err != nil {
			fmt.Printf("%s : %v\n", email, err)
		} else {
			fmt.Printf("%s: Space[%v] Get[%v] Put[%v] Transfer[%v] Price[%v]\n", email, godata.FriendlyByte(order.Space), order.ApicallGet, order.ApicallPut, godata.FriendlyByte(order.Transfer), order.Price)
		}
	} else {
		fmt.Printf("%s skip...\n", email)
	}
}

func monthstartstop(value, resultformat string) (string, string) {
	timer, _ := time.Parse(gotime.Y+gotime.M, value)
	year := timer.Year()
	month := timer.Month()
	days := 0
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return fmt.Sprintf(resultformat, year, month, "01"), fmt.Sprintf(resultformat, year, month, days)
}
