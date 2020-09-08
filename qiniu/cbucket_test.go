package qiniu

import ()

const (
	TEST_AK = "Ti7caH9lslWoNGRlPv6LIvwP1iiCTE2I4pszkHig"
	TEST_SK = "YX2iSwJfwftz_KWz3zcEh1XHF9vRqbMxyuMNSYVC"
)

/*
func Test_KeyCreateBucket(t *testing.T){
	gohttp.REQUEST_DEBUG=true
	req := KeyCreateBucketRequest{}
	req.AccessKey=TEST_AK
	req.SecretKey=TEST_SK
	req.Name="qianfan-test"
	req.IsPublic=BUCKET_PUBLIC
	resp := &KeyCreateBucketResponse{}
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
*/

/*
func Test_KeyDeleteBucket(t *testing.T){
		req :=KeyDeleteBucketRequest{}
		req.AccessKey=TEST_AK
		req.SecretKey=TEST_SK
		req.Name="qianfan-test"
		resp:=&KeyDeleteBucketResponse{}
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
*/

/*
func Test_KeyListBucket(t *testing.T) {
	gohttp.REQUEST_DEBUG=true
	access, _ := getqiniuaccess("2789378367@qq.com", "19850314@qq.com")
	child, _ := getqiniuchild(1380436258, *access)
	req := KeyListBucketRequest{}
	req.AccessKey=child.Key
	req.SecretKey=child.Secret
	fmt.Println(child.Key)
	fmt.Println(child.Secret)
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
			}else {
				t.Logf("%v", resp)
			}
		}else {
			//object
			resp := &KeyListBucketDomainResponse{}
			err := json.Unmarshal(bytes, resp)
			if err != nil {
				t.Errorf("%v", err)
			}else {
				t.Errorf("%v", resp.Msg())
			}
		}
	}
}
*/

/*
func Test_KeyListBucketDomain(t *testing.T) {
	req := KeyListBucketDomainRequest{}
	req.BucketName="qianfan-test"
	req.AccessKey=TEST_AK
	req.SecretKey=TEST_SK
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
			}else {
				t.Logf("%v", resp)
			}
		}else {
			//object
			resp := &KeyListBucketDomainResponse{}
			err := json.Unmarshal(bytes, resp)
			if err != nil {
				t.Errorf("%v", err)
			}else {
				t.Errorf("%v", resp.Msg())
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

func getqiniuchild(uid int64, access Access) (Child, error) {
	child := Child{}
	child.Uid = uid
	path := GetDir() + string(os.PathSeparator) + child.GetFileName()
	if gofile.FileExists(path) {
		err := child.ReadKeys(true, access.AccessToken)
		if err != nil {
			return Child{}, err
		}
	}
	if len(child.Key) == 0 || len(child.Secret) == 0 || child.IsExpired() {
		err := child.Get(access.AccessToken)
		if err != nil {
			return Child{}, err
		}
	}
	return child, nil
}

func Test_ChildStatInfo(t *testing.T) {
	gohttp.REQUEST_DEBUG=true
	//access, _ := getqiniuaccess("2789378367@qq.com", "19850314@qq.com")
	//child, _ := getqiniuchild(1380436258, *access)
	req := &ChildStatInfoRequest{}
	//req.AccessKey=child.Key
	//req.SecretKey=child.Secret
	req.AccessKey="d4klLjig9JqQUZEZhBin5yft2oCepqm21ci7OjeI"
	req.SecretKey="Ti30_N96ZdOS47-oB4PcghzPUpQy_xjx8rsG46t0"
	req.Uid=1380436258
	req.Bucket="qianfanyun-18qiang"
	req.Month="201509"
	resp := &ChildStatInfoResponse{}
	err := gohttp.DoGetResponse(req, resp)
	if err != nil {
		t.Error(err)
	}else {
		t.Logf("%v\n", resp)
	}
}
*/
