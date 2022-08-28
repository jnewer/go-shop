package order

import (
	"github.com/gin-gonic/gin"
	"go-shop/domain/order"
	"go-shop/utils/api_helper"
	"go-shop/utils/pagination"
	"net/http"
)

type Controller struct {
	orderService *order.Service
}

func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{orderService: orderService}
}

// CompleteOrder godoc
// @Summary 完成订单
// @Tags    Order
// @Accept  json
// @Produce json
// @Param   Authorization header   string true "Authentication header"
// @Success 200           {object} api_helper.Response
// @Failure 400           {object} api_helper.ErrorResponse
// @Router  /order [post]
func (c *Controller) Complete(g *gin.Context) {
	userId := api_helper.GetUserId(g)
	err := c.orderService.Complete(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, api_helper.Response{
		Message: "Order Created",
	})
}

// Cancel godoc
// @Summary 取消订单
// @Tags    Order
// @Accept  json
// @Produce json
// @Param   Authorization      header   string             true "Authentication header"
// @Param   CancelOrderRequest body     CancelOrderRequest true "order information"
// @Success 200                {object} api_helper.Response
// @Failure 400                {object} api_helper.ErrorResponse
// @Router  /order [delete]
func (c *Controller) Cancel(g *gin.Context) {
	var req CancelOrderRequest

	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)
	err := c.orderService.Cancel(userId, req.OrderId)

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, api_helper.Response{Message: "Order Canceled"})
}

// GetAll godoc
// @Summary 获得订单列表
// @Tags    Order
// @Accept  json
// @Produce json
// @Param   Authorization header   string true  "Authentication header"
// @Param   page          query    int    false "Page number"
// @Param   pageSize      query    int    false "Page size"
// @Success 200           {object} pagination.Pages
// @Router  /order [get]
func (c *Controller) GetAll(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userId := api_helper.GetUserId(g)
	page = c.orderService.GetAll(page, userId)
	g.JSON(http.StatusOK, page)
}
