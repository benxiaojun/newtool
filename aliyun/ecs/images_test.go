package ecs

import (
	_ "encoding/json"
	_ "git.sillydong.com/chenzhidong/goczd/gohttp"
	_ "testing"
)

/*
func Test_DescribeImage(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&DescribeImagesRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.RegionId="cn-hangzhou"
	req.ImageOwnerAlias="self"
	resp := &DescribeImagesResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		if resp.OK() {
			//str, err := json.MarshalIndent(resp, "", "    ")
			//if err != nil {
			//	t.Errorf("%v", err)
			//}else {
			//	t.Log(string(str))
			//}
			for _,image := range resp.Images.Image{
				t.Logf("%v",image.ImageId)
			}
		}else {
			t.Errorf("%v", resp.Msg())
		}
	}

}
*/
