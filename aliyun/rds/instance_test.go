package rds

import ()

/*
func Test_DescribeDBInstances(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &DescribeDBInstancesRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"
	req.PageSize=100

	resp:=&DescribeDBInstancesResponse{}
	err:=gohttp.DoGetResponse(req,resp)
	if err != nil {
		t.Errorf("%v",err)
	}else{
		if resp.OK() {
			for _,dbinstance := range resp.Items.DBInstance{
				t.Logf("%s %s",dbinstance.DBInstanceDescription,dbinstance.ExpireTime)
			}
		}else {
			t.Errorf("%v", resp.Msg())
		}
	}
}
*/

/*
func Test_DescribeDbInstanceAttribute(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &DescribeDBInstanceAttributeRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rds9nh9vbgs7ftuqfscl5"

	resp := &DescribeDBInstanceAttributeResponse{}
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
func Test_RestartDBInstance(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &RestartDBInstanceRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"

	resp := &RestartDBInstanceResponse{}
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
func Test_ModifySecurityIps(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&ModifySecurityIpsRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.SecurityIps="10.252.168.79"

	resp:=&ModifySecurityIpsResponse{}
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
func Test_UpgradeDBInstanceEngine(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&UpgradeDBInstanceEngineVersionRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.EngineVersion="5.6"

	resp:=&UpgradeDBInstanceEngineVersionResponse{}
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
func Test_ModifyDBInstanceDescription(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&ModifyDBInstanceDescriptionRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rds9nh9vbgs7ftuqfscl5"
	req.DBInstanceDescription="化龙巷"
	resp:=&ModifyDBInstanceDescriptionResponse{}
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
