package repositories

import (
	"github.com/chadiso/golang_tryout/app/models"
	"github.com/jinzhu/gorm"
)

func GetTransactions(conn *gorm.DB) models.Transactions {
	var transactions models.Transactions

	conn.Find(&transactions)

	return transactions
}

func GetTransaction(conn *gorm.DB, transactionID string) models.Transaction {
	var transaction models.Transaction

	conn.First(&transaction, transactionID)

	return transaction
}
