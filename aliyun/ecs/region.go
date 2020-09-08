package ecs

import ()

type DescribeRegionsResponse struct {
	EcsBaseResponse
	Regions struct {
		Region []RegionType `json:"Region"`
	} `json:"Regions"`
}

type DescribeRegionsRequest struct {
	EcsBaseRequest
}

func (req *DescribeRegionsRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeRegions"
	req.Child = req
	return req.Sign()
}

type DescribeZonesResponse struct {
	EcsBaseResponse
	Zones struct {
		Zone []ZoneType `json:"Zone"`
	} `json:"Zones"`
}

type DescribeZonesRequest struct {
	EcsBaseRequest
	RegionId string `url:"RegionId"`
}

func (req *DescribeZonesRequest) B() error {
	if err := req.EcsBaseRequest.B(); err != nil {
		return err
	}
	req.Action = "DescribeZones"
	req.Child = req
	return req.Sign()
}
