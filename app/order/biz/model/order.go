package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Name          string `gorm:"type:varchar(255)"`
	Email         string `gorm:"type:varchar(255)"`
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint32      `gorm:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) (orders []*Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).Preload("OrderItems").Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return
}
