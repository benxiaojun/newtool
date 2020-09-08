package qiniu

import ()

/*
func Test_CreateChildAccount(t *testing.T){
	access := Access{}
	err := access.ReadOauth()
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := CreateChildAccountRequest{}
		req.AuthorizationValue=access.AccessToken
		req.Email="test@qianfanyun.com"
		req.Password="test4qianfanyun"
		resp := &CreateChildAccountResponse{}
		err:=gohttp.DoPostResponse(&req,resp)
		if err != nil {
			t.Errorf("%v",err)
		}else{
			if resp.OK(){
				t.Logf("%v",resp)
			}else{
				t.Logf("%v",resp.Msg())
			}
		}
	}
}
*/

/*
func Test_ListChildAccount(t *testing.T){
	access := Access{}
	err := access.ReadOauth(true)
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := ListChildAccountRequest{}
		req.AuthorizationValue=access.AccessToken

		bytes,err := gohttp.DoGet(&req)
		if err != nil {
			t.Errorf("%v\n",err)
		}else{
			if string(bytes[0])=="[" {
				//array
				resp:=[]ListChildAccountResponse{}
				err := json.Unmarshal(bytes, &resp)
				if err != nil {
					t.Errorf("%v", err)
				}else {
					t.Logf("%v", resp)
				}
			}else{
				//object
				resp:=&ListChildAccountResponse{}
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
func Test_DisableChildAccount(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	access := Access{}
	err := access.ReadOauth()
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := DisableChildAccountRequest{}
		req.AuthorizationValue=access.AccessToken
		req.Uid=1380485649//test@qianfanyun.com
		req.Reason="测试"

		resp:=&DisableChildAccountResponse{}
		err:=gohttp.DoPostResponse(&req,resp)
		if err != nil {
			t.Errorf("%v",err)
		}else{
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
func Test_EnableChildAccount(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	access := Access{}
	err := access.ReadOauth()
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := EnableChildAccountRequest{}
		req.AuthorizationValue=access.AccessToken
		req.Uid=1380485649//test@qianfanyun.com

		resp := &EnableChildAccountResponse{}
		err := gohttp.DoPostResponse(&req, resp)
		if err != nil {
			t.Errorf("%v", err)
		}else {
			if resp.OK() {
				t.Logf("%v", resp)
			}else {
				t.Errorf("%v", resp.Msg())
			}
		}
	}
}
*/

/*
func Test_ChildKeys(t *testing.T) {
	gohttp.REQUEST_DEBUG=true

	access := Access{}
	err := access.ReadOauth()
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := ChildKeysRequest{}
		req.AuthorizationValue=access.AccessToken
		req.Uid=1380485649 //test@qianfanyun.com
		resp := &ChildKeysResponse{}
		err := gohttp.DoGetResponse(&req, resp)
		if err != nil {
			t.Errorf("%v", err)
		}else {
			if resp.OK() {
				t.Logf("%v", resp)
			}else {
				t.Errorf("%v", resp.Msg())
			}
		}
	}
}
*/
