package qiniu

import ()

/*
func Test_WriteOauth(t *testing.T){
	access := Access{}
	access.AccessToken="a"
	access.ExpiresIn=3600
	access.RefreshToken="b"
	err:=access.WriteOauth()
	if err != nil {
		t.Errorf("%v\n",err)
	} else {
		t.Log("WriteOauth 测试通过")
	}
}
*/

/*
func Test_ReadOauth(t *testing.T){
	access:=Access{}
	access.Username="285273540@qq.com"
	err:=access.ReadOauth(false)
	if err != nil {
		t.Errorf("%v\n",err)
	} else {
		if access.AccessToken=="a" && access.ExpiresIn==3600 && access.RefreshToken=="b"{
			t.Log("ReadOauth 测试通过")
		} else {
			t.Errorf("%v\n",access)
		}
	}
}
*/

/*
func Test_Get(t *testing.T) {
	access := Access{}
	access.Username="285273540@qq.com"
	err := access.Get("285273540@qq.com", "nanjibing@9989")
	if err!=nil {
		t.Errorf("%v\n", err)
	}else {
		if access.AccessToken!="" && access.ExpiresIn>0 && access.RefreshToken!="" {
			t.Log("正确")
		} else {
			t.Errorf("%v\n", access)
		}
	}
}
*/

/*
func Test_Refresh(t *testing.T) {
	access := Access{}
	err := access.ReadOauth()
	if err != nil {
		t.Errorf("%v\n", err)
	} else {
		oldaccess:=access
		err:=access.Refresh(true)
		if err != nil {
			t.Errorf("%v\n",err)
		}else if access!=oldaccess{
			t.Logf("%v\n",access)
		}else{
			t.Errorf("刷新失败:%v",access)
		}
	}
}
*/

/*
type TestRequest struct {
	QiniuBaseRequest
	Uid int `url:"-"`
}

func (req *TestRequest)URL() string {
	return "http://rs.qiniu.com/move/bmV3ZG9jczpmaW5kX21hbi50eHQ=/bmV3ZG9jczpmaW5kLm1hbi50eHQ=";
}

func Test_Sign(t *testing.T){
	req:=TestRequest{}
	accesskey="MY_ACCESS_KEY"
	secretkey="MY_SECRET_KEY"
	sign,err:=Sign(&req,accesskey,secretkey)
	if err != nil {
		t.Errorf("%v",err)
	}else{
		t.Log(sign)
	}

}
*/
