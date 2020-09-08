package rds

import ()

/*
func Test_CreateAccount(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &CreateAccountRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.AccountName="taccout"
	req.AccountPassword="chenzzd00123"
	req.AccountDescription="testapi"

	resp :=&CreateAccountResponse{}
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
func Test_DeleteAccount(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &DeleteAccountRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.AccountName="taccout"

	resp:=&DeleteAccountResponse{}
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
func Test_GrantAccountPrivilege(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&GrantAccountPrivilegeRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.AccountName="taccout"
	req.DBName="taccout"
	req.AccountPrivilege=ACCOUNT_PRIVILEGE_READWRITE

	resp:=&GrantAccountPrivilegeResponse{}
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
func Test_RevokeAccountPrivilege(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	req:=&RevokeAccountPrivilegeRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.AccountName="taccout"
	req.DBName="taccout"

	resp:=&RevokeAccountPrivilegeResponse{}
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
func Test_DescribeAccounts(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&DescribeAccountsRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.AccountName="taccout"

	resp :=&DescribeAccountsResponse{}
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
func Test_ModifyAccountDescription(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &ModifyAccountDescriptionRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.AccountName="taccout"
	req.AccountDescription="hello"

	resp := &ModifyAccountDescriptionResponse{}
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
