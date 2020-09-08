package rds

import ()

/*
func Test_CreateDatabase(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&CreateDatabaseRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.DBName="taccout"
	req.CharacterSetName="utf8"
	req.DBDescription="taccout"

	resp:=&CreateDatabaseResponse{}
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
func Test_DeleteDatabase(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&DeleteDatabaseRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.DBName="taccout"

	resp:=&DeleteDatabaseResponse{}
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
func Test_DescribeDatabases(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req:=&DescribeDatabasesRequest{}
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.DBName="ssp"
	//req.DBStatus=DB_STATUS_RUNNING
	resp:=&DescribeDatabasesResponse{}
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
func Test_ModifyDBDescription(t *testing.T){
	gohttp.REQUEST_DEBUG=true

	req := &ModifyDBDescriptionRequest{};
	req.AccessKeyId=ACCESS_KEY_ID
	req.AccessSecret=ACCESS_KEY_SECRET
	req.DBInstanceId="rdsltmqi8nd7880tf4z4"
	req.DBName="taccout"
	req.DBDescription="hello"

	resp:=&ModifyDBDescriptionResponse{}
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
