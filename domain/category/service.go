package category

import (
	"go-shop/utils/csv_helper"
	"go-shop/utils/pagination"
	"mime/multipart"
)

type Service struct {
	r Repository
}

func NewCategoryService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()

	return &Service{
		r: r,
	}
}

func (s *Service) Create(category *Category) error {
	existCate := s.r.GetByName(category.Name)
	if len(existCate) > 0 {
		return ErrCategoryExistWithName
	}

	err := s.r.Create(category)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, nil
	}

	for _, categoryVariable := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariable[0], categoryVariable[1]))
	}

	count, err := s.r.BulkCreate(categories)

	if err != nil {
		return count, err
	}
	return count, nil
}

func (s *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := s.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}
