package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
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

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("Orderitems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
