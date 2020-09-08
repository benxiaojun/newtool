package rds

import ()

/*
func Test_DescribeSlowLogs(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&DescribeSlowLogsRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rds9nh9vbgs7ftuqfscl5"
	req.StartTime="2015-08-15Z"
	req.EndTime="2015-08-19Z"

	resp:=&DescribeSlowLogsResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		if resp.OK() {
			str, err := json.MarshalIndent(resp, "", "    ")
			if err != nil {
				t.Errorf("%v", err)
			}else {
				t.Log(string(str))
			}
		}else {
			t.Errorf("%v", resp.Msg())
		}
	}
}
*/

/*
func Test_DescribeErrorLogs(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &DescribeErrorLogsRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rds9nh9vbgs7ftuqfscl5"
	req.StartTime="2015-07-15T00:00Z"
	req.EndTime="2015-08-18T00:00Z"

	resp := &DescribeErrorLogsResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		if resp.OK() {
			str, err := json.MarshalIndent(resp, "", "    ")
			if err != nil {
				t.Errorf("%v", err)
			}else {
				t.Log(string(str))
			}
		}else {
			t.Errorf("%v", resp.Msg())
		}
	}
}
*/
