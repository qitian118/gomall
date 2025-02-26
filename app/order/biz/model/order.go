package model

import "gorm.io/gorm"

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);unique_index"`
	UserId       uint32      `gorm:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	Orderitems   []OrderItem `gorm:"foreignKey:OrderIdRefer;reference:OrderId"`
}

func (Order) TableName() string {
	return "order"
}
