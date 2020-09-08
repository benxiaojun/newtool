package rds

import "time"

type DescribeSlowLogsResponse struct {
	RdsBaseResponse
	EndTime string `json:"EndTime"`
	Engine  string `json:"Engine"`
	Items   struct {
		SQLSlowLog []struct {
			CreateTime                string `json:"CreateTime"`
			DBName                    string `json:"DBName"`
			MaxExecutionTime          int64  `json:"MaxExecutionTime"`
			MaxLockTime               int64  `json:"MaxLockTime"`
			MySQLTotalExecutionCounts int64  `json:"MySQLTotalExecutionCounts"`
			MySQLTotalExecutionTimes  int64  `json:"MySQLTotalExecutionTimes"`
			ParseMaxRowCount          int64  `json:"ParseMaxRowCount"`
			ParseTotalRowCounts       int64  `json:"ParseTotalRowCounts"`
			ReportTime                string `json:"ReportTime"`
			ReturnMaxRowCount         int64  `json:"ReturnMaxRowCount"`
			ReturnTotalRowCounts      int64  `json:"ReturnTotalRowCounts"`
			SQLText                   string `json:"SQLText"`
			TotalLockTimes            int64  `json:"TotalLockTimes"`
		} `json:"SQLSlowLog"`
	} `json:"Items"`
	PageNumber       int64  `json:"PageNumber"`
	PageRecordCount  int64  `json:"PageRecordCount"`
	StartTime        string `json:"StartTime"`
	TotalRecordCount int64  `json:"TotalRecordCount"`
}

type DescribeSlowLogsRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	StartTime    string `url:"StartTime"`
	EndTime      string `url:"EndTime"`
	DBName       string `url:"DBName,omitempty"`
	SortKey      string `url:"SortKey,omitempty"`
	PageSize     int    `url:"PageSize,omitempty"`
	PageNumber   int    `url:"PageNumber,omitempty"`
}

func (req *DescribeSlowLogsRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeSlowLogs"
	req.Child = req
	return req.Sign()
}

type DescribeErrorLogsResponse struct {
	RdsBaseResponse
	Items struct {
		ErrorLog []struct {
			CreateTime time.Time `json:"CreateTime"`
			ErrorInfo  string    `json:"ErrorInfo"`
		} `json:"ErrorLog"`
	} `json:"Items"`
	PageNumber       int64 `json:"PageNumber"`
	PageRecordCount  int64 `json:"PageRecordCount"`
	TotalRecordCount int64 `json:"TotalRecordCount"`
}

type DescribeErrorLogsRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	StartTime    string `url:"StartTime"`
	EndTime      string `url:"EndTime"`
	PageSize     int    `url:"PageSize,omitempty"`
	PageNumber   int    `url:"PageNumber,omitempty"`
}

func (req *DescribeErrorLogsRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeErrorLogs"
	req.Child = req
	return req.Sign()
}
