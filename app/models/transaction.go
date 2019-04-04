package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Transaction struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type Transactions []*Transaction

func (Transaction) TableName() string {
	return "transactions"
}
