package order

import (
	"go-shop/domain/cart"
	"go-shop/domain/product"
	"go-shop/utils/pagination"
	"time"
)

var day14Hours float64 = 336

type Service struct {
	orderRepository       Repository
	orderedItemRepository OrderItemRepository
	productRepository     product.Repository
	cartRepository        cart.Repository
	cartItemRepository    cart.ItemRepository
}

func NewService(
	orderRepository Repository,
	orderedItemRepository OrderItemRepository,
	productRepository product.Repository,
	cartRepository cart.Repository,
	cartItemRepository cart.ItemRepository) *Service {
	orderRepository.Migration()
	orderedItemRepository.Migration()
	return &Service{
		orderRepository:       orderRepository,
		orderedItemRepository: orderedItemRepository,
		productRepository:     productRepository,
		cartRepository:        cartRepository,
		cartItemRepository:    cartItemRepository}
}

func (s *Service) Complete(userId uint) error {
	currentCart, err := s.cartRepository.FindOrCreateByUserID(userId)

	if err != nil {
		return err
	}

	cartItems, err := s.cartItemRepository.GetItems(currentCart.UserID)
	if err != nil {
		return err
	}

	if len(cartItems) == 0 {
		return ErrEmptyCartFound
	}

	orderedItems := make([]OrderedItem, 0)

	for _, item := range cartItems {
		orderedItems = append(orderedItems, *NewOrderedItem(item.Count, item.ProductID))
	}

	err = s.orderRepository.Create(NewOrder(userId, orderedItems))
	return err
}

func (s *Service) Cancel(uid, oid uint) error {
	currentOrder, err := s.orderRepository.FindByOrderID(oid)

	if err != nil {
		return err
	}

	if currentOrder.UserID != uid {
		return ErrInvalidOrderID
	}

	if currentOrder.CreatedAt.Sub(time.Now()).Hours() > day14Hours {
		return ErrCancelDurationPassed
	}

	currentOrder.IsCanceled = true
	err = s.orderRepository.Update(*currentOrder)

	return err
}

func (s *Service) GetAll(page *pagination.Pages, uid uint) *pagination.Pages {
	orders, count := s.orderRepository.GetAll(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count

	return page
}
