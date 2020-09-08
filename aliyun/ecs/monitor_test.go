package ecs

import ()

/*
func Test_DescribeInstanceMonitorData(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &DescribeInstanceMonitorDataRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.InstanceId="i-233hq61pz"
	req.SetStartTime("2015-08-03 00:00:00")
	req.SetEndTime("2015-08-03 01:00:00")
	req.Period=PERIOD_ONE_MINUTE
	resp :=&DescribeInstanceMonitorDataResponse{}
	err:=gohttp.DoGetResponse(req,resp)
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
