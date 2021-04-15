package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

type StockIn struct {
	gorm.Model
	Time              time.Time `gorm:"not_null"`
	ProductId         int       `gorm:"not_null"`
	Product           Product
	PricePerProduct   float32 `gorm:"not_null"`
	TransactionNumber string  `gorm:"unique_index;not_null"`
	Note              string
	OrderedQuantity   int
	ReceivedQuantity  int
}
