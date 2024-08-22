package dao

import (
	"gorm.io/gorm"
)

var (
	_dao *DAO
)

// DAO data access object.
type DAO struct {
	db *gorm.DB
}

// Init base dao.
func Init(db *gorm.DB) {
	_dao = &DAO{
		db: db,
	}
}

// Get base dao.
func Get() *DAO {
	return _dao
}

func (dao *DAO) WithTransaction() (*TxDAO, error) {
	tx := dao.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &TxDAO{DAO: DAO{db: tx}}, nil
}
