package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-shop/domain/category"
	"go-shop/utils/api_helper"
	"go-shop/utils/pagination"
	"net/http"
)

type Controller struct {
	categoryService *category.Service
}

func NewCategoryController(categoryService *category.Service) *Controller {
	return &Controller{categoryService: categoryService}
}

// Create godoc
// @Summary 根据给定的参数创建分类
// @Tags    Category
// @Accept  json
// @Produce json
// @Param   Authorization         header   string                true "Authentication header"
// @Param   CreateCategoryRequest body     CreateCategoryRequest true "category information"
// @Success 200                   {object} api_helper.Response
// @Failure 400                   {object} api_helper.ErrorResponse
// @Router  /category [post]
func (c *Controller) Create(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	newCategory := category.NewCategory(req.Name, req.Desc)
	err := c.categoryService.Create(newCategory)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSONP(http.StatusCreated, api_helper.Response{Message: "Category created"})
}

// BulkCreate godoc
// @Summary 根据给定的csv文件，批量创建分类
// @Tags    Category
// @Accept  json
// @Produce json
// @Param   Authorization header   string true "Authentication header"
// @Param   file          formData file   true "file contains category information"
// @Success 200           {object} api_helper.Response
// @Failure 400           {object} api_helper.ErrorResponse
// @Router  /category/upload [post]
func (c Controller) BulkCreate(g *gin.Context) {
	fileHeader, err := g.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	count, err := c.categoryService.BulkCreate(fileHeader)

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSONP(http.StatusOK, api_helper.Response{Message: fmt.Sprintf("'%s' uploaded: '%d' new categories created", fileHeader.Filename, count)})
}

// GetAll godoc
// @Summary 获得分类列表
// @Tags    Category
// @Accept  json
// @Produce json
// @Param   page     query    int false "Page number"
// @Param   pageSize query    int false "Page size"
// @Success 200      {object} pagination.Pages
// @Router  /category [get]
func (c *Controller) GetAll(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.categoryService.GetAll(page)
	g.JSON(http.StatusOK, page)

}
