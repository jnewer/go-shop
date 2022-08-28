package product

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		log.Print(err)
	}
}
func (r *Repository) FindBySKU(sku string) (*Product, error) {
	var product *Product
	err := r.db.Where("IsDeleted=?", 0).Where(Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

func (r *Repository) Update(updateProduct Product) error {
	savedProduct, err := r.FindBySKU(updateProduct.SKU)
	if err != nil {
		return err
	}
	err = r.db.Model(&savedProduct).Updates(updateProduct).Error
	return err
}

func (r *Repository) SearchByString(str string, pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	convertedStr := "%" + str + "%"
	var count int64
	r.db.Where("IsDeleted=?", false).Where(
		"Name like ? or SKU like ?", convertedStr,
		convertedStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

func (r *Repository) Create(p *Product) error {
	result := r.db.Create(p)

	return result.Error
}

func (r *Repository) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64

	r.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

func (r *Repository) Delete(sku string) error {
	currentProduct, err := r.FindBySKU(sku)
	if err != nil {
		return err
	}

	currentProduct.IsDeleted = true

	return r.db.Save(currentProduct).Error
}
