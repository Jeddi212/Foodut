package model

import (
	"time"

	prModel "github.com/Foodut/backend/modules/product/domain/model"
)

type Transaction struct {
	ID              int               `form:"id" json:"id" gorm:"primaryKey"`
	CustomerID      int               `form:"customerId" json:"customerId"`
	Address         string            `form:"address" json:"address"`
	PaymentOption   string            `form:"paymentOption" json:"paymentOption"`
	SubTotal        float64           `form:"subTotal" json:"subTotal"`
	TransactionDate time.Time         `form:"transactionDate" json:"transactionDate"`
	ProductDetail   []prModel.Product `form:"productDetail" json:"productDetail" gorm:"many2many:transaction_details"`
}
