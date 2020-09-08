package rds

import (
	"fmt"
	"github.com/benxiaojun/satool/aliyun/rds"
	"github.com/sillydong/goczd/gofile"
	"github.com/sillydong/goczd/gohttp"
	"github.com/sillydong/goczd/gotime"
)

func Listinstances(id, secret string) {
	req := &rds.DescribeDBInstancesRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.RegionId = "cn-hangzhou"
	req.PageSize = 100
	req.PageNumber = 1
	resp := &rds.DescribeDBInstancesResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		if resp.OK() {
			if len(resp.Items.DBInstance) > 0 {
				for index, instance := range resp.Items.DBInstance {
					fmt.Printf("[%03d]   %s[%s]\n", index, instance.DBInstanceID, instance.DBInstanceDescription)
					fmt.Printf("        %s.mysql.rds.aliyuncs.com\n", instance.DBInstanceID)
					fmt.Printf("        [%s]%s %.1f\n", instance.DBInstanceStatus, instance.Engine, instance.EngineVersion)
				}
			} else {
				fmt.Print("此账号无实例记录\n")
			}
		} else {
			fmt.Errorf("Error: %s\n", resp.Msg())
		}
	}
}

func Listslowlog(id, secret string) {
	var rdsid, start, stop string
	fmt.Print("RDS实例ID:")
	fmt.Scanf("%s\n", &rdsid)

	fmt.Print("日志开始时间(例:2015-01-01):")
	fmt.Scanf("%s\n", &start)
	fmt.Print("日志结束时间(例:2015-01-01):")
	fmt.Scanf("%s\n", &stop)

	if len(rdsid) > 0 && len(start) > 0 && len(stop) > 0 {
		req := &rds.DescribeSlowLogsRequest{}
		req.AccessKeyId = id
		req.AccessSecret = secret
		req.DBInstanceId = rdsid
		req.StartTime = start + "Z"
		req.EndTime = stop + "Z"

		resp := &rds.DescribeSlowLogsResponse{}
		err := gohttp.DoGetResponse(req, resp)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			if resp.OK() {
				if len(resp.Items.SQLSlowLog) > 0 {
					var content string
					content += fmt.Sprintf("%s\n%s - %s\n\n", rdsid, start, stop)
					for index, log := range resp.Items.SQLSlowLog {
						content += fmt.Sprintf("%s\n%s\n最长执行时间: %d 最长锁时间: %d 处理行数: %d 返回结果行数: %d 执行次数: %d\n报告时间: %s\n\n", log.DBName, log.SQLText, log.MaxExecutionTime, log.MaxLockTime, log.ParseMaxRowCount, log.ReturnMaxRowCount, log.MySQLTotalExecutionCounts, log.CreateTime)
						fmt.Printf("[%04d]   数据库:%s\n", index, log.DBName)
						fmt.Printf("         %s\n", log.SQLText)
						fmt.Printf("         最长执行时间: %d 最长锁时间: %d 处理行数: %d 返回结果行数: %d 执行次数: %d\n", log.MaxExecutionTime, log.MaxLockTime, log.ParseMaxRowCount, log.ReturnMaxRowCount, log.MySQLTotalExecutionCounts)
						fmt.Printf("         报告时间: %s\n", log.CreateTime)
					}
					gofile.IoUtilWriteFile(rdsid+"_slow_"+gotime.GetTimeStr(gotime.Y+gotime.M+gotime.D+gotime.H+gotime.I+gotime.S)+".txt", []byte(content), 777)

				} else {
					fmt.Print("无慢查询日志\n")
				}
			} else {
				fmt.Errorf("%s\n", resp.Msg())
			}
		}
	} else {
		fmt.Print("输入无效\n")
	}
}

func Listerrorlog(id, secret string) {
	var rdsid, start, stop string
	fmt.Print("RDS实例ID:")
	fmt.Scanf("%s\n", &rdsid)

	fmt.Print("日志开始时间(例:2015-01-01):")
	fmt.Scanf("%s\n", &start)
	fmt.Print("日志结束时间(例:2015-01-01):")
	fmt.Scanf("%s\n", &stop)

	if len(rdsid) > 0 && len(start) > 0 && len(stop) > 0 {
		req := &rds.DescribeErrorLogsRequest{}
		req.AccessKeyId = id
		req.AccessSecret = secret
		req.DBInstanceId = rdsid
		req.StartTime = start + "T00:00Z"
		req.EndTime = stop + "T23:59Z"

		resp := &rds.DescribeErrorLogsResponse{}
		err := gohttp.DoGetResponse(req, resp)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			if resp.OK() {
				if len(resp.Items.ErrorLog) > 0 {
					var content string
					content += fmt.Sprintf("%s\n%s - %s\n\n", rdsid, start, stop)
					for _, log := range resp.Items.ErrorLog {
						content += fmt.Sprintf("%s\n%s\n\n", log.CreateTime, log.ErrorInfo)
						fmt.Printf("[%s]%s\n", log.CreateTime, log.ErrorInfo)
					}
					gofile.IoUtilWriteFile(rdsid+"_error_"+gotime.GetTimeStr(gotime.Y+gotime.M+gotime.D+gotime.H+gotime.I+gotime.S)+".txt", []byte(content), 777)
				} else {
					fmt.Print("无错误日志\n")
				}
			} else {
				fmt.Errorf("%s\n", resp.Msg())
			}
		}
	} else {
		fmt.Print("输入无效\n")
	}
}

func Export(id, secret, kind string) {
	var start, stop string

	fmt.Print("日志开始时间(例:2015-01-01):")
	fmt.Scanf("%s\n", &start)
	fmt.Print("日志结束时间(例:2015-01-01):")
	fmt.Scanf("%s\n", &stop)

	req := &rds.DescribeDBInstancesRequest{}
	req.AccessKeyId = id
	req.AccessSecret = secret
	req.RegionId = "cn-hangzhou"
	req.PageSize = 100
	req.PageNumber = 1
	resp := &rds.DescribeDBInstancesResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		if resp.OK() {
			if len(resp.Items.DBInstance) > 0 {
				time := gotime.GetTimeStr(gotime.Y + gotime.M + gotime.D + gotime.H + gotime.I + gotime.S)

				for _, instance := range resp.Items.DBInstance {
					fmt.Println("Getting: " + instance.DBInstanceID)

					switch kind {
					case "slow":
						req := &rds.DescribeSlowLogsRequest{}
						req.AccessKeyId = id
						req.AccessSecret = secret
						req.DBInstanceId = instance.DBInstanceID
						req.StartTime = start + "Z"
						req.EndTime = stop + "Z"

						resp := &rds.DescribeSlowLogsResponse{}
						err := gohttp.DoGetResponse(req, resp)
						if err != nil {
							fmt.Printf("%v\n", err)
						} else {
							if resp.OK() {
								if len(resp.Items.SQLSlowLog) > 0 {
									var content string
									content += fmt.Sprintf("%s(%s)\n%s - %s\n\n", instance.DBInstanceID, instance.DBInstanceDescription, start, stop)
									for _, log := range resp.Items.SQLSlowLog {
										content += fmt.Sprintf("%s\n%s\n最长执行时间: %d 最长锁时间: %d 处理行数: %d 返回结果行数: %d 执行次数: %d\n报告时间: %s\n\n", log.DBName, log.SQLText, log.MaxExecutionTime, log.MaxLockTime, log.ParseMaxRowCount, log.ReturnMaxRowCount, log.MySQLTotalExecutionCounts, log.CreateTime)
									}
									if err := gofile.OsWriteFile("slow_"+time+".txt", []byte(content), 0644, true); err != nil {
										fmt.Printf("%s\n", err)
									}

								} else {
									fmt.Print("无慢查询日志\n")
								}
							} else {
								fmt.Errorf("%s\n", resp.Msg())
							}
						}
					case "error":
						req := &rds.DescribeErrorLogsRequest{}
						req.AccessKeyId = id
						req.AccessSecret = secret
						req.DBInstanceId = instance.DBInstanceID
						req.StartTime = start + "T00:00Z"
						req.EndTime = stop + "T23:59Z"

						resp := &rds.DescribeErrorLogsResponse{}
						err := gohttp.DoGetResponse(req, resp)
						if err != nil {
							fmt.Printf("%v\n", err)
						} else {
							if resp.OK() {
								if len(resp.Items.ErrorLog) > 0 {
									var content string
									content += fmt.Sprintf("%s(%s)\n%s - %s\n\n", instance.DBInstanceID, instance.DBInstanceDescription, start, stop)
									for _, log := range resp.Items.ErrorLog {
										content += fmt.Sprintf("%s\n%s\n\n", log.CreateTime, log.ErrorInfo)
									}
									if err := gofile.OsWriteFile("error_"+time+".txt", []byte(content), 0644, true); err != nil {
										fmt.Printf("%v\n", err)
									}
								} else {
									fmt.Print("无错误日志\n")
								}
							} else {
								fmt.Errorf("%s\n", resp.Msg())
							}
						}
					}

				}
			} else {
				fmt.Print("此账号无实例记录\n")
			}
		} else {
			fmt.Errorf("Error: %s\n", resp.Msg())
		}
	}
}
