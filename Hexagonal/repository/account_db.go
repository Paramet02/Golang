package repository

import (
	"gorm.io/gorm"
)

type accountRpeositoryDB struct {
	db *gorm.DB
}

func NewaccountRpeositoryDB(db *gorm.DB) AccountRpeository {
	return &accountRpeositoryDB{db : db}
}

func (a accountRpeositoryDB) Create(acc Account) (*Account, error) {
	// ใช้ gorm ในการบันทึกข้อมูล
	if err := a.db.Create(&acc).Error; err != nil {
		return nil, err
	}

	// gorm จะอัพเดท AccountID โดยอัตโนมัติ ถ้าตารางมี auto-increment
	return &acc, nil
}
func (a accountRpeositoryDB) GetAll(id int) ([]Account , error) {
	accounts := []Account{}

	// SELECT * FROM accounts WHERE customer_id = ?;
	if result := a.db.Where("customer_id = ?", id).Find(&accounts); result.Error != nil {
		return nil, result.Error
	}

	return accounts , nil
}