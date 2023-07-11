package mysql

import (
	"gorm.io/gorm"

	"pearshop_backend/app/domain/repository"
	"pearshop_backend/pkg/gormutil"
)

// TxDataSQL manages the transaction by implementing the Transaction Manager interface
type TxDataSQL struct {
	DB *gorm.DB
}

// NewTxDataSQL is the constructor function
func NewTxDataSQL() repository.TransactionManager {
	return &TxDataSQL{}
}

// TxBegin begin a new transaction
func (tds *TxDataSQL) TxBegin() {
	db := gormutil.GetDB()
	tds.DB = db.Begin()
}

// TxRollback rollback a transaction
func (tds *TxDataSQL) TxRollback() {
	tds.DB.Rollback()
}

// TxCommit commit a transaction
func (tds *TxDataSQL) TxCommit() error {
	return tds.DB.Commit().Error
}

// GetTx to get the current transaction of this service
func (tds *TxDataSQL) GetTx() interface{} {
	return tds.DB
}
