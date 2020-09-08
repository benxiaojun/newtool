package newqiniu

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/sillydong/goczd/gohttp"
)

type Accounts struct {
	Type    int
	Company int
	UID     int
	UserID  string
	Email   string
	Month   string
	AK      string
	SK      string
}

const (
	EMAIL    = "2789378367@qq.com"
	PASSWORD = "19850314@qq.com"
)

func TestQiniu(t *testing.T) {
	gohttp.REQUEST_DEBUG = true
	// 获取管理凭证
	req := &AccountOauthTokenRequest{
		Username: EMAIL,
		Password: PASSWORD,
	}

	resp := &AccountOauthTokenResponse{}
	err := gohttp.DoPostResponse(req, resp)
	if err != nil {
		t.Error(err)
	} else {
		if resp.OK() {
			fmt.Printf("%+v\n", resp)
			inforeq := &AccountInfoRequest{}
			inforeq.QiniuBaseRequest.AuthorizationValue = resp.AccessToken
			inforesp := &AccountInfoResponse{}
			err := gohttp.DoPostResponse(inforeq, inforesp)
			if err != nil {
				t.Error(err)
			} else {
				if inforesp.OK() {
					fmt.Printf("%+v\n", inforesp)
					kreq := &AccountKeysRequest{
						App: "default",
					}
					kreq.QiniuBaseRequest.AuthorizationValue = resp.AccessToken
					kresp := &AccountKeysResponse{}
					err := gohttp.DoPostResponse(kreq, kresp)
					if err != nil {
						t.Error(err)
					} else {
						if kresp.OK() {
							fmt.Printf("%+v\n", kresp)

							lreq := &ListChildAccountRequest{
								Offset: 0,
								Limit:  1000,
							}
							lreq.AuthorizationValue = resp.AccessToken
							// lreq.AccessKey = kresp.Key
							// lreq.SecretKey = kresp.Secret
							lresp := &ListChildAccountResponse{}
							err := gohttp.DoPostResponse(lreq, lresp)
							if err != nil {
								t.Error(err)
							} else {
								fmt.Printf("%+v\n", lresp)
							}

							// ckreq := &ChildKeysRequest{
							// 	Uid: 1380436258,
							// }
							// ckreq.AuthorizationValue = resp.AccessToken
							// // ckreq.AccessKey = kresp.Key
							// // ckreq.SecretKey = kresp.Secret
							// ckresp := &ChildKeysResponse{}
							// err := gohttp.DoPostResponse(ckreq, ckresp)
							// if err != nil {
							// 	t.Error(err)
							// } else {
							// 	fmt.Printf("%+v\n", ckresp)
							// }

							breq := &ListBucketRequest{}
							breq.AccessKey = kresp.Key
							breq.SecretKey = kresp.Secret
							// breq.AuthorizationValue = resp.AccessToken
							// breq.QiniuBaseRequest.AuthorizationValue = resp.AccessToken

							bytes, err := gohttp.DoPost(breq)
							if err != nil {
								t.Error(err)
							} else {
								if string(bytes[0]) == "[" {
									var bresp []string
									err := json.Unmarshal(bytes, &bresp)
									if err != nil {
										t.Error(err)
									} else {
										fmt.Printf("%+v\n", bresp)

										var domains []string
										// var totalspace int64

										for _, bucket := range bresp {
											// bsreq := &BucketStatInfoRequest{
											// 	Uid:    inforesp.UID,
											// 	Bucket: bucket,
											// 	Month:  "201605",
											// }
											// bsreq.AccessKey = kresp.Key
											// bsreq.SecretKey = kresp.Secret
											// bsresp := &BucketStatInfoResponse{}
											// err := gohttp.DoGetResponse(bsreq, bsresp)
											// if err != nil {
											// 	t.Error(err)
											// } else {
											// 	fmt.Printf("%+v\n", bsresp)
											// 	totalspace += bsresp.Space
											// }

											dreq := &ListBucketDomainRequest{}
											dreq.BucketName = bucket
											dreq.AccessKey = kresp.Key
											dreq.SecretKey = kresp.Secret
											dbytes, err := gohttp.DoGet(dreq)
											if err != nil {
												t.Error(err)
											} else {
												if string(dbytes[0]) == "[" {
													var dresp []string
													err := json.Unmarshal(dbytes, &dresp)
													if err != nil {
														t.Error(err)
													} else {
														fmt.Printf("%+v\n", dresp)
														domains = append(domains, dresp...)
													}
												} else {
													t.Error(string(dbytes))
												}
											}
										}
										if len(domains) > 0 {
											treq := &DomainTrafficRequest{
												Domains:     strings.Join(domains, ";"),
												Granularity: GRANULARITY_DAY,
												StartDate:   "2016-07-01",
												EndDate:     "2016-07-31",
											}
											treq.AccessKey = kresp.Key
											treq.SecretKey = kresp.Secret
											tresp := &DomainTrafficResponse{}
											err := gohttp.DoPostJsonResponse(treq, tresp)
											if err != nil {
												t.Error(err)
											} else {
												if tresp.OK() {
													fmt.Printf("%+v\n", tresp)
												} else {
													t.Error(tresp.Msg())
												}
											}
										}

										// fmt.Printf("totalspace: %v\n", totalspace)

										// preq := &UsagePutRequest{
										// 	Begin:       "20160601/00:00",
										// 	End:         "20160701/00:00",
										// 	Granularity: GRANULARITY_DAY,
										// }
										// preq.AccessKey = kresp.Key
										// preq.SecretKey = kresp.Secret
										// pbytes, err := gohttp.DoGet(preq)
										// if err != nil {
										// 	t.Error(err)
										// } else {
										// 	pjson, err := simplejson.NewJson(pbytes)
										// 	if err != nil {
										// 		t.Error(pjson)
										// 	} else {
										// 		if pjson.IsObject() {
										// 			t.Error(pjson.Get("error").MustString())
										// 		} else {
										// 			var stotalput int64
										// 			for _, item := range pjson.MustArray() {
										// 				hits, _ := (item.(map[string]interface{})["values"]).(map[string]interface{})["hits"].(json.Number).Int64()
										// 				stotalput += hits
										// 				// stotalput += ((item.(map[string]interface{})["values"]).(map[string]interface{})["hits"]).(int64)
										// 				// fmt.Printf("%t\n%v\n", item, item)
										// 			}
										// 			fmt.Printf("totalput %v\n", stotalput)
										// 		}
										// 	}

										// 	// if err != nil {
										// 	// 	t.Error(err)
										// 	// } else {
										// 	// 	if _, ok := pjson.CheckGet("error"); ok {
										// 	// 		t.Error(pjson.Get("error").MustString())
										// 	// 	} else {
										// 	// 		m, err := pjson.Array()
										// 	// 		if err != nil {
										// 	// 			t.Error(err)
										// 	// 		} else {
										// 	// 			for k, v := range m {
										// 	// 				fmt.Printf("%v ------ %v\n", k, v)
										// 	// 			}
										// 	// 		}
										// 	// 	}
										// 	// }
										// }

										// sreq := &UsageSpaceRequest{
										// 	Begin:       "20160731000000",
										// 	End:         "20160801000000",
										// 	Granularity: GRANULARITY_DAY,
										// }
										// sreq.AccessKey = kresp.Key
										// sreq.SecretKey = kresp.Secret
										// sresp := &UsageSpaceResponse{}
										// err2 := gohttp.DoGetResponse(sreq, sresp)
										// if err2 != nil {
										// 	t.Error(err2)
										// } else {
										// 	if sresp.OK() {
										// 		fmt.Printf("%+v\n", sresp.Datas)
										// 		var stotalspace int64
										// 		for _, num := range sresp.Datas {
										// 			stotalspace += num
										// 		}
										// 		fmt.Printf("xtotalspace: %+v\n", stotalspace)
										// 	} else {
										// 		t.Error(sresp.Msg())
										// 	}
										// }
									}
								} else {
									bresp := &ListBucketResponse{}
									err := json.Unmarshal(bytes, bresp)
									if err != nil {
										t.Error(err)
									} else {
										if resp.OK() {
											fmt.Printf("%+v\n", bresp)
										} else {
											t.Error(resp.Msg())
										}
									}
								}
							}
						} else {
							t.Error(kresp.Msg())
						}
					}
				} else {
					t.Error(inforesp.Msg())
				}
			}
		} else {
			t.Error(resp.Msg())
		}
	}
}
