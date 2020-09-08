package ecs

import ()

/*
func Test_DescribeInstances(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:= &DescribeInstancesRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"
	//bytes,_:=json.Marshal([]string{"i-233hq61pz"})
	//req.InstanceIds=string(bytes)

	resp:=&DescribeInstancesResponse{}
	err:=gohttp.DoGetResponse(req,resp)
	if err != nil {
		t.Errorf("%v",err)
	}else{
		if resp.OK(){

			//str,err:=json.MarshalIndent(resp,"","    ")
			//if err != nil {
			//	t.Errorf("%v",err)
			//}else{
			//	t.Log(string(str))
			//}

			for _,instance := range resp.Instances.Instance{
				t.Logf("%v\n",instance.ImageID)
			}
		}else{
			t.Errorf("%v",resp.Msg())
		}
	}

}
*/

/*
func Test_DescribeInstanceStatus(t *testing.T) {
	gohttp.REQUEST_DEBUG=true

	req := &DescribeInstanceStatusRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"
	req.ZoneId="cn-hangzhou-d"

	resp := &DescribeInstanceStatusResponse{}
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
func Test_CreateInstance(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&CreateInstanceRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"
	req.ZoneId="cn-hangzhou-d"
	req.ImageId="centos6u5_64_20G_aliaegis_20150130.vhd"//待定
	//req.InstanceType="ecs.s2.large"//2核4G
	req.InstanceType="ecs.s3.large"//4核8G
	req.SecurityGroupId="sg-23zyz1oog"
	req.InstanceName="数字蚌埠"
	req.Description=""
	req.InternetChargeType="PayByBandwidth"
	//req.InternetMaxBandwidthIn="-1"
	req.InternetMaxBandwidthOut="5"
	req.Password="1LokMAppnlFQ"
	//req.IoOptimized="optimized"

	resp:=&CreateInstanceResponse{}
	err:=gohttp.DoGetResponse(req,resp)
	if err != nil {
		t.Errorf("%v",err)
	}else{
		if resp.OK() {
			str,err:=json.MarshalIndent(resp,"","    ")
			if err != nil {
				t.Errorf("%v",err)
			}else{
				t.Log(string(str))
			}
		}else {
			t.Errorf("%v", resp.Msg())
		}
	}
}
*/

/*
func Test_ModifyInstanceAttribute(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&ModifyInstanceAttributeRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.InstanceId="i-23uksrrm1"
	req.InstanceName="化龙巷"
	req.Description=""
	resp:=&ModifyInstanceAttributeResponse{}
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
