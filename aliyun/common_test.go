package aliyun

import ()

/*
type TestResponse struct{
	AliyunBaseResponse
}

type TestRequest struct{
	AliyunBaseRequest
	RegionId string `url:"RegionId"`
}

func Test_Sign(t *testing.T){
	req := TestRequest{}
	req.Action="DescribeDBInstances"
	req.AccessKeyId="testid"
	req.AccessSecret="testsecret"
	req.RegionId="region1"
	req.B()
	sign,err :=Sign(&req,req.AccessSecret)
	if err != nil {
		t.Errorf("%v",err)
	}else{
		t.Log(sign)
	}

}
*/
