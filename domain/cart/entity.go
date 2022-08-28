package cart

import (
	"go-shop/domain/product"
	"go-shop/domain/user"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `json:"foreignKey:ID;references:UserID"`
}

//NewCart 实例化
func NewCart(userID uint) *Cart {
	return &Cart{UserID: userID}
}

type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

//NewCartItem 实例化
func NewCartItem(productID uint, cartID uint, count int) *Item {
	return &Item{ProductID: productID, Count: count, CartID: cartID}
}
