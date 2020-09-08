package ecs

//查询可用镜像返回数据结构
type DescribeImagesResponse struct {
	EcsBaseResponse
	RegionId   string `json:"RegionId"`
	TotalCount int64  `json:"TotalCount"`
	PageNumber int64  `json:"PageNumber"`
	PageSize   int64  `json:"PageSize"`
	Images     struct {
		Image []ImageType `json:"Image"`
	} `json:"Images"`
}

//查询可用镜像请求数据结构
type DescribeImagesRequest struct {
	EcsBaseRequest
	RegionId        string `url:"RegionId"`
	ImageId         string `url:"ImageId,omitempty"`
	SnapshotId      string `url:"SnapshotId,omitempty"`
	ImageName       string `url:"ImageName,omitempty"`
	ImageOwnerAlias string `url:"ImageOwnerAlias,omitempty"`
	PageNumber      int64  `url:"PageNumber,omitempty"`
	PageSize        int64  `url:"PageSize,omitempty"`
}

//填充默认数据，签名
func (req *DescribeImagesRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeImages"
	req.Child = req
	return req.Sign()
}
