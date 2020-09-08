package ecs

import ()

/*
func Test_DescribeSecurityGroups(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &DescribeSecurityGroupsRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"
	resp := &DescribeSecurityGroupsResponse{}
	err:=gohttp.DoGetResponse(req,resp)
	if err != nil {
		t.Errorf("%v",err)
	}else{
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
func Test_DescribeSecurityGroupAttribute(t *testing.T) {
	gohttp.REQUEST_DEBUG=true

	req := &DescribeSecurityGroupAttributeRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"
	req.SecurityGroupId="sg-23zyz1oog"
	resp := &DescribeSecurityGroupAttributeResponse{}
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
