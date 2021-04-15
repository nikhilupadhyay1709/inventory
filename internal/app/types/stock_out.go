package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

type StockOut struct {
	gorm.Model
	Time            time.Time `gorm:"not_null"`
	ProductId       int       `gorm:"not_null"`
	Product         Product
	PricePerProduct float32 `gorm:"not_null"`
	Note            string
	Quantity        int
}
