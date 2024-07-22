package models

import (
	"time"

	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	ECDSA   string  `gorm:"type:varchar(255)" json:"ECDSA" binding:"required"`
	EDDSA   string  `gorm:"type:varchar(255)" json:"EdDSA" binding:"required"`
	Chain   string  `json:"chain"`
	Address string  `json:"address"`
	Token   string  `json:"token"`
	Balance float64 `json:"balance"`
	Date    int64   `json:"date"`
	PriceID uint    `json:"price_id"`
}

type BalanceResponse struct {
	ID         uint      `json:"id"`
	ECDSA      string    `json:"ecdsa"`
	EDDSA      string    `json:"eddsa"`
	Chain      string    `json:"chain"`
	Address    string    `json:"address"`
	Token      string    `json:"token"`
	Balance    float64   `json:"balance"`
	Date       int64     `json:"date"`
	PriceID    uint      `json:"price_id"`
	UsdBalance float64   `json:"usd_balance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Price      Price     `json:"price"`
}
