package utils

import "gorm.io/gorm"

func Paginate(db *gorm.DB, page int, pageSize int, result interface{}) error {
	offset := (page - 1) * pageSize
	err := db.Limit(pageSize).Offset(offset).Find(result).Error
	if err != nil {
		return err
	}
	return nil
}
