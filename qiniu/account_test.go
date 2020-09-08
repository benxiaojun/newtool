package qiniu

import ()

/*
func Test_AccountInfo(t *testing.T){
	access, _ := getqiniuaccess("2789378367@qq.com", "19850314@qq.com")
	req := AccountInfoRequest{}
	req.AuthorizationValue=access.AccessToken
	resp := &AccountInfoResponse{}
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
*/

/*
func Test_AccountKeys(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	access := Access{}
	err := access.ReadOauth()
	if err != nil {
		t.Errorf("%v", err)
	}else {
		req := AccountKeysRequest{}
		req.AuthorizationValue=access.AccessToken
		req.App="default"
		resp:=&AccountKeysResponse{}
		err:=gohttp.DoGetResponse(&req,resp)
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
func getqiniuaccess(username, password string) (*Access, error) {
	access := &Access{}
	access.Username = username
	if gofile.FileExists(GetDir() + string(os.PathSeparator) + access.GetFileName()) {
		err := access.ReadOauth(true)
		if err != nil {
			return nil, err
		}
	}
	if len(access.AccessToken) == 0 || access.IsExpired() {
		err := access.Get(username, password)
		if err != nil {
			return nil, err
		}
	}
	return access, nil
}
*/

/*
func Test_AccountStat(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	access,_ := getqiniuaccess("2789378367@qq.com","19850314@qq.com")
	req := AccountStatInfoRequest{}
	req.AuthorizationValue=access.AccessToken
	req.Uid=1380434442
	req.Month="201508"
	resp := &AccountStatInfoResponse{}
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
*/
