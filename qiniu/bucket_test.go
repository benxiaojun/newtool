package qiniu

import ()

/*
func Test_CreateBucket(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	access := Access{}
	err := access.ReadOauth(true)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := CreateBucketRequest{}
		req.Name="qianfan-test"
		req.IsPublic=BUCKET_PUBLIC
		req.AuthorizationValue=access.AccessToken
		resp := &CreateBucketResponse{}
		err:=gohttp.DoPostResponse(&req,resp)
		if err != nil {
			t.Errorf("%v",err)
		}else {
			if resp.OK(){
				t.Logf("%v",resp)
			}else{
				t.Errorf("%v",resp.Msg())
			}
		}
	}
}
*/

/*
func Test_DeleteBucket(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	access := Access{}
	err := access.ReadOauth(true)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req :=DeleteBucketRequest{}
		req.Name="qianfan-test"
		req.AuthorizationValue=access.AccessToken
		resp:=&DeleteBucketResponse{}
		err:=gohttp.DoPostResponse(&req,resp)
		if err != nil {
			t.Errorf("%v",err)
		}else{
			if resp.OK(){
				t.Log("%v",resp)
			}else{
				t.Errorf("%v",resp.Msg())
			}
		}
	}
}
*/

/*
func Test_ListBucket(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	access := Access{}
	err := access.ReadOauth(true)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req:=ListBucketRequest{}
		req.AuthorizationValue=access.AccessToken
		bytes, err := gohttp.DoPost(&req)
		if err != nil {
			t.Errorf("%v", err)
		}else {
			if string(bytes[0])=="[" {
				//array
				resp := make([]string, 0)
				err := json.Unmarshal(bytes, &resp)
				if err != nil {
					t.Errorf("%v", err)
				}else{
					t.Logf("%v",resp)
				}
			}else {
				//object
				resp := &ListBucketDomainResponse{}
				err := json.Unmarshal(bytes, resp)
				if err != nil {
					t.Errorf("%v", err)
				}else {
					t.Errorf("%v", resp.Msg())
				}
			}
		}
	}
}
*/

/*
func Test_ListBucketDomain(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	access := Access{}
	err := access.ReadOauth(true)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := ListBucketDomainRequest{}
		req.BucketName="qianfan-test"
		req.AuthorizationValue=access.AccessToken
		bytes,err:=gohttp.DoPost(&req)
		if err != nil {
			t.Errorf("%v",err)
		}else{
			if string(bytes[0])=="["{
				//array
				resp:=make([]string,0)
				err := json.Unmarshal(bytes,&resp)
				if err != nil {
					t.Errorf("%v",err)
				}else{
					t.Logf("%v",resp)
				}
			}else{
				//object
				resp:=&ListBucketDomainResponse{}
				err:=json.Unmarshal(bytes,resp)
				if err != nil {
					t.Errorf("%v",err)
				}else{
					t.Errorf("%v",resp.Msg())
				}
			}
		}
	}
}
*/
