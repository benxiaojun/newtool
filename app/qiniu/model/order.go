package model

import (
	"fmt"
	"github.com/sillydong/goczd/gotime"
)

const (
	TYPE_ACCOUNT = 1
	TYPE_CHILD   = 2
)
const (
	COMPANY_QIANFAN = 1
	COMPANY_HANGJIA = 2
)

func CreateOrder(order *Order) error {
	isExists, err := OrderExists(order.Company, order.Email, order.Month)
	if err != nil {
		return err
	} else if isExists {
		return fmt.Errorf("[%s]%s exists", order.Month, order.Email)
	}

	order.Importtime = int(gotime.GetTimestamp())
	order.Paytime = 0
	order.Invoicetime = 0
	if _, err = x.InsertOne(order); err != nil {
		return err
	}
	return nil
}

func DropMonth(company int, month string) (int64, error) {
	return x.Delete(&Order{Company: company, Month: month})
}

func OrderExists(company int, email, month string) (bool, error) {
	return x.Get(&Order{Company: company, Email: email, Month: month})
}

func GetOrders(company int, email, month string) ([]Order, error) {
	orders := make([]Order, 0)
	query := x.Where("1=1")
	query.And("Company = ?", company)
	if len(email) > 0 {
		query.And("email = ?", email)
	}
	if len(month) > 0 {
		query.And("month = ?", month)
	}
	if err := query.Asc("uid").Find(&orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func SumOrders(company int, email, month string) ([]map[string][]byte, error) {
	sql := "select sum(price) as sum_price,sum(space) as sum_space,sum(apicall_get) as sum_apicall_get, sum(apicall_put) as sum_apicall_put,sum(transfer) as sum_transfer from `order`"
	if len(email) > 0 {
		return x.Query(sql+" where company = ? and email = ?", company, email)
	} else if len(month) > 0 {
		return x.Query(sql+" where company = ? and month = ?", company, month)
	} else {
		return nil, fmt.Errorf("无筛选条件")
	}
}
