package ecs

import ()

/*
func Test_DescribeRegions(t *testing.T){
	//gohttp.REQUEST_DEBUG=true

	req := &DescribeRegionsRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET

	resp:=&DescribeRegionsResponse{}
	err:=gohttp.DoGetResponse(req,resp)
	if err != nil {
		t.Errorf("%v",err)
	}else{
		if resp.OK(){
			str,err:=json.MarshalIndent(resp,"","    ")
			if err != nil {
				t.Errorf("%v",err)
			}else {
				t.Log(string(str))
			}
		}else{
			t.Errorf("%v",resp.Msg())
		}
	}
}
*/

/*
func Test_DescribeZones(t *testing.T) {
	gohttp.REQUEST_DEBUG=true

	req := &DescribeZonesRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"

	resp := &DescribeZonesResponse{}
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
