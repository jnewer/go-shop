package order

import (
	"gorm.io/gorm"
	"log"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) Migration() {
	err := r.db.AutoMigrate(&OrderedItem{})

	if err != nil {
		log.Print(err)
	}
}

func (r *OrderItemRepository) Update(item OrderedItem) error {
	result := r.db.Save(&item)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderItemRepository) Create(ci *OrderedItem) error {
	result := r.db.Create(ci)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
