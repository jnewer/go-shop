package product

import "go-shop/utils/pagination"

type Service struct {
	productRepository Repository
}

//NewService 构造函数
func NewService(productRepository Repository) *Service {
	productRepository.Migration()
	return &Service{productRepository: productRepository}
}

//GetAll 获取所有商品分页
func (s *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count := s.productRepository.GetAll(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count

	return page
}

//Create 创建商品
func (s *Service) Create(name string, desc string, count int, price float32, cid uint) error {
	newProduct := NewProduct(name, desc, count, price, cid)
	err := s.productRepository.Create(newProduct)

	return err
}

//Delete 删除商品
func (s *Service) Delete(sku string) error {
	err := s.productRepository.Delete(sku)

	return err
}

//Update 更新商品
func (s *Service) Update(product *Product) error {
	err := s.productRepository.Update(*product)

	return err
}

//Search 查找商品
func (s *Service) Search(text string, page *pagination.Pages) *pagination.Pages {
	products, count := s.productRepository.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}
