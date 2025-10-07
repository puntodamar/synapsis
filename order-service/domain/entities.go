package domain

import "time"

type Order struct {
	ID         string      `json:"id" gorm:"primaryKey;type:varchar(255)"`
	CustomerID string      `json:"customer_id" gorm:"not null;type:varchar(255)"`
	Status     Status      `json:"status" gorm:"type:varchar(32);not null;index"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID      uint   `json:"-" gorm:"primaryKey;autoIncrement"`
	OrderID string `json:"-" gorm:"not null;type:varchar(255);index:uniq_order_sku,unique"`
	SKU     string `json:"sku" gorm:"not null;type:varchar(255);index:uniq_order_sku,unique"`
	Qty     int32  `json:"qty" gorm:"not null"`
}
