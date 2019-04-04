package repositories

import (
	"github.com/chadiso/golang_tryout/app/models"
	"github.com/jinzhu/gorm"
)

func GetTransactions(conn *gorm.DB) models.Transactions {
	transactions := models.Transactions{}

	conn.Find(&transactions)

	return transactions
}
