package cart

import (
	"github.com/gin-gonic/gin"
	"go-shop/domain/cart"
	"go-shop/utils/api_helper"
	"net/http"
)

type Controller struct {
	cartService *cart.Service
}

func NewCartController(cartService *cart.Service) *Controller {
	return &Controller{cartService: cartService}
}

// AddItem godoc
// @Summary 添加Item
// @Tags    Cart
// @Accept  json
// @Produce json
// @Param   Authorization   header   string          true "Authentication header"
// @Param   ItemCartRequest body     ItemCartRequest true "product information"
// @Success 200             {object} api_helper.Response
// @Failure 400             {object} api_helper.ErrorResponse
// @Router  /cart/item [post]
func (c *Controller) AddItem(g *gin.Context) {
	var req ItemCartRequest

	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	userId := api_helper.GetUserId(g)
	err := c.cartService.AddItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSONP(http.StatusCreated, CreateCategoryResponse{Message: "Item added to cart"})
}

// UpdateItem godoc
// @Summary 更新Item
// @Tags    Cart
// @Accept  json
// @Produce json
// @Param   Authorization   header   string          true "Authentication header"
// @Param   ItemCartRequest body     ItemCartRequest true "product information"
// @Success 200             {object} api_helper.Response
// @Failure 400             {object} api_helper.ErrorResponse
// @Router  /cart/item [patch]
func (c *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	userId := api_helper.GetUserId(g)

	err := c.cartService.UpdateItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, CreateCategoryResponse{Message: "updated"})
}

func (c *Controller) GetCart(g *gin.Context) {
	userId := api_helper.GetUserId(g)
	result, err := c.cartService.GetCartItems(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, result)
}
