package model

type Order struct {
	Id                  int     `xorm:"not null pk autoincr INT(11)"`
	Company             int     `xorm:"not null TINYINT(1)"`
	Type                int     `xorm:"not null TINYINT(1)"`
	Uid                 int     `xorm:"not null index INT(11)"`
	Userid              string  `xorm:"not null VARCHAR(64)"`
	Email               string  `xorm:"not null index VARCHAR(64)"`
	Month               string  `xorm:"not null index VARCHAR(10)"`
	Price               string  `xorm:"not null DECIMAL(32,4)"`
	Space               int64   `xorm:"not null default 0 BIGINT(32)"`
	SpaceAvg            int64   `xorm:"not null default 0 BIGINT(32)"`
	ApicallGet          int64   `xorm:"not null default 0 BIGINT(32)"`
	ApicallPut          int64   `xorm:"not null default 0 BIGINT(32)"`
	Transfer            int64   `xorm:"not null default 0 BIGINT(32)"`
	PriceSpaceUnit      int     `xorm:"not null default 0 INT(11)"`
	PriceSpace          float64 `xorm:"not null default 0.0000 DECIMAL(10,4)"`
	PriceApicallGetUnit int     `xorm:"not null default 0 INT(11)"`
	PriceApicallGet     float64 `xorm:"not null default 0.0000 DECIMAL(10,4)"`
	PriceApicallPutUnit int     `xorm:"not null default 0 INT(11)"`
	PriceApicallPut     float64 `xorm:"not null default 0.0000 DECIMAL(10,4)"`
	PriceTransferUnit   int     `xorm:"not null default 0 INT(11)"`
	PriceTransfer       float64 `xorm:"not null default 0.0000 DECIMAL(10,4)"`
	Importtime          int     `xorm:"not null default 0 INT(11)"`
	Paytime             int     `xorm:"not null default 0 INT(11)"`
	Invoicetime         int     `xorm:"not null default 0 INT(11)"`
}
