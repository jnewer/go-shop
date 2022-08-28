package cart

import "gorm.io/gorm"

func (item *Item) AfterUpdate(tx *gorm.DB) (err error) {
	if item.Count <= 0 {
		return tx.Unscoped().Delete(&item).Error
	}

	return
}
