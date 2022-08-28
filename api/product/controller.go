package product

import (
	"github.com/gin-gonic/gin"
	"go-shop/domain/product"
	"go-shop/utils/api_helper"
	"go-shop/utils/pagination"
	"net/http"
)

type Controller struct {
	productService product.Service
}

func NewProductController(productService product.Service) *Controller {
	return &Controller{productService: productService}
}

// GetProducts godoc
// @Summary 获得商品列表（分页）
// @Tags    Product
// @Accept  json
// @Produce json
// @Param   qt       query    string false "Search text to find matched sku numbers and names"
// @Param   page     query    int    false "Page number"
// @Param   pageSize query    int    false "Page size"
// @Success 200      {object} pagination.Pages
// @Router  /product [get]
func (c *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	queryText := g.Query("qt")
	if queryText != "" {
		page = c.productService.Search(queryText, page)
	} else {
		page = c.productService.GetAll(page)
	}

	g.JSON(http.StatusOK, page)
}

// Create godoc
// @Summary 创建商品
// @Tags    Product
// @Accept  json
// @Produce json
// @Param   Authorization        header   string               true "Authentication header"
// @Param   CreateProductRequest body     CreateProductRequest true "product information"
// @Success 200                  {object} api_helper.Response
// @Failure 400                  {object} api_helper.ErrorResponse
// @Router  /product [post]
func (c *Controller) Create(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.Create(req.Name, req.Desc, req.Count, req.Price, req.CategoryID)

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusCreated, api_helper.Response{Message: "Product Created"})
}

// Delete godoc
// @Summary 删除商品根据sku
// @Tags    Product
// @Accept  json
// @Produce json
// @Param   DeleteProductRequest body     DeleteProductRequest true "sku of product"
// @Param   Authorization        header   string               true "Authentication header"
// @Success 200                  {object} api_helper.Response
// @Failure 400                  {object} api_helper.ErrorResponse
// @Router  /product [delete]
func (c *Controller) Delete(g *gin.Context) {
	var req DeleteProductRequest

	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.Delete(req.SKU)

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, api_helper.Response{Message: "Product Deleted"})
}

// Update godoc
// @Summary 更新商品更加sku
// @Tags    Product
// @Accept  json
// @Produce json
// @Param   UpdateProductRequest body     UpdateProductRequest true "product information"
// @Param   Authorization        header   string               true "Authentication header"
// @Success 200                  {object} CreateProductResponse
// @Failure 400                  {object} api_helper.ErrorResponse
// @Router  /product [patch]
func (c *Controller) Update(g *gin.Context) {
	var req UpdateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.Update(req.ToProduct())

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, CreateProductResponse{Message: "Product Updated"})
}
