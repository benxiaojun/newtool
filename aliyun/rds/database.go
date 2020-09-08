package rds

type CreateDatabaseResponse struct {
	RdsBaseResponse
}

type CreateDatabaseRequest struct {
	RdsBaseRequest
	DBInstanceId     string `url:"DBInstanceId"`
	DBName           string `url:"DBName"`
	CharacterSetName string `url:"CharacterSetName"`
	DBDescription    string `url:"DBDescription"`
}

func (req *CreateDatabaseRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "CreateDatabase"
	req.Child = req
	return req.Sign()
}

type DeleteDatabaseResponse struct {
	RdsBaseResponse
}

type DeleteDatabaseRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	DBName       string `url:"DBName"`
}

func (req *DeleteDatabaseRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DeleteDatabase"
	req.Child = req
	return req.Sign()
}

type DescribeDatabasesResponse struct {
	RdsBaseResponse
	Databases struct {
		DataBase []DataBase `json:"Database"`
	} `json:"Databases"`
}

type DescribeDatabasesRequest struct {
	RdsBaseRequest
	DBInstanceId string `url:"DBInstanceId"`
	DBName       string `url:"DBName,omitempty"`
	DBStatus     string `url:"DBStatus,omitempty"`
}

func (req *DescribeDatabasesRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeDatabases"
	req.Child = req
	return req.Sign()
}

type ModifyDBDescriptionResponse struct {
	RdsBaseResponse
}

type ModifyDBDescriptionRequest struct {
	RdsBaseRequest
	DBInstanceId  string `url:"DBInstanceId"`
	DBName        string `url:"DBName"`
	DBDescription string `url:"DBDescription"`
}

func (req *ModifyDBDescriptionRequest) B() error {
	if err := req.RdsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "ModifyDBDescription"
	req.Child = req
	return req.Sign()
}
