package cart

import (
	"errors"
	"go-shop/domain/product"
)

type Service struct {
	cartRepository     Repository
	cartItemRepository ItemRepository
	productRepository  product.Repository
}

func NewService(cartRepository Repository, cartItemRepository ItemRepository, productRepository product.Repository) *Service {
	cartRepository.Migration()
	cartItemRepository.Migration()
	return &Service{cartRepository: cartRepository, cartItemRepository: cartItemRepository, productRepository: productRepository}
}

func (s *Service) AddItem(userID uint, sku string, count int) error {
	currentProduct, err := s.productRepository.FindBySKU(sku)

	if err != nil {
		return err
	}

	currentCart, err := s.cartRepository.FindOrCreateByUserID(userID)
	if err != nil {
		return err
	}
	_, err = s.cartItemRepository.FindByID(currentProduct.ID, currentCart.ID)

	if err != nil {
		return ErrItemAlreadyExistInCart
	}

	if currentProduct.StockCount < count {
		return product.ErrProductStockIsNotEnough
	}

	if count <= 0 {
		return ErrCountInvalid
	}

	err = s.cartItemRepository.Create(NewCartItem(currentProduct.ID, currentCart.ID, count))

	return err

}

func (s *Service) UpdateItem(userID uint, sku string, count int) error {
	currentProduct, err := s.productRepository.FindBySKU(sku)
	if err != nil {
		return err
	}

	currentCart, err := s.cartRepository.FindByUserID(userID)

	if err != nil {
		return err
	}

	currentItem, err := s.cartItemRepository.FindByID(currentProduct.ID, currentCart.ID)
	if err != nil {
		return errors.New("item 不存在")
	}

	if currentProduct.StockCount+currentItem.Count < count {
		return product.ErrProductStockIsNotEnough
	}

	currentItem.Count = count

	err = s.cartItemRepository.Update(*currentItem)

	return err
}

func (s *Service) GetCartItems(userID uint) ([]Item, error) {
	currentCart, err := s.cartRepository.FindOrCreateByUserID(userID)

	if err != nil {
		return nil, err
	}

	items, err := s.cartItemRepository.GetItems(currentCart.ID)

	if err != nil {
		return nil, err
	}

	return items, nil
}
