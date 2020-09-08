package rds

type CreateAccountResponse struct {
	RdsBaseResponse
}

type CreateAccountRequest struct {
	RdsBaseRequest
	DBInstanceId       string `url:"DBInstanceId"`
	AccountName        string `url:"AccountName"`
	AccountPassword    string `url:"AccountPassword"`
	AccountDescription string `url:"AccountDescription"`
}

func (req *CreateAccountRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "CreateAccount"
	req.Child = req
	return req.Sign()
}

type DeleteAccountResponse struct {
	RdsBaseResponse
}

type DeleteAccountRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	AccountName  string `url:"AccountName"`
}

func (req *DeleteAccountRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DeleteAccount"
	req.Child = req
	return req.Sign()
}

type GrantAccountPrivilegeResponse struct {
	RdsBaseResponse
}

type GrantAccountPrivilegeRequest struct {
	RdsBaseRequest
	DBInstanceId     string `url:"DBInstanceId"`
	AccountName      string `url:"AccountName"`
	DBName           string `url:"DBName"`
	AccountPrivilege string `url:"AccountPrivilege"`
}

func (req *GrantAccountPrivilegeRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "GrantAccountPrivilege"
	req.Child = req
	return req.Sign()
}

type RevokeAccountPrivilegeResponse struct {
	RdsBaseResponse
}

type RevokeAccountPrivilegeRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	AccountName  string `url:"AccountName"`
	DBName       string `url:"DBName"`
}

func (req *RevokeAccountPrivilegeRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "RevokeAccountPrivilege"
	req.Child = req
	return req.Sign()
}

type DescribeAccountsResponse struct {
	RdsBaseResponse
	Accounts struct {
		DBInstanceAccount []DBInstanceAccount `json:"DBInstanceAccount"`
	} `json:"Accounts"`
}

type DescribeAccountsRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	AccountName  string `url:"AccountName"`
}

func (req *DescribeAccountsRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeAccounts"
	req.Child = req
	return req.Sign()
}

type ModifyAccountDescriptionResponse struct {
	RdsBaseResponse
}

type ModifyAccountDescriptionRequest struct {
	RdsBaseRequest
	DBInstanceId       string `url:"DBInstanceId"`
	AccountName        string `url:"AccountName"`
	AccountDescription string `url:"AccountDescription"`
}

func (req *ModifyAccountDescriptionRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "ModifyAccountDescription"
	req.Child = req
	return req.Sign()
}
