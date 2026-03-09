package handlers

import (
	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"com.DotNicolasPenha.SistemaDePedidos/internal/httpg/responses"
	"com.DotNicolasPenha.SistemaDePedidos/internal/service"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service service.OrderService
}

func NewOrderHandler(g *gin.Engine, orderService service.OrderService) {
	ordersRoute := g.Group("/orders")
	orderHandler := orderHandler{service: orderService}
	{
		ordersRoute.GET("/", orderHandler.getOrders)
		ordersRoute.GET("/:id", orderHandler.getOneOrderById)
		ordersRoute.POST("/", orderHandler.postOrder)
	}
}

func (h *orderHandler) getOrders(ctx *gin.Context) {
	orders, err := h.service.ListAll()
	if err != nil {
		responses.ReponseInternalError(ctx, err)
		return
	}
	responses.ResponseOkData(ctx, orders)
}
func (h *orderHandler) postOrder(ctx *gin.Context) {
	var OrderDto domain.CreateOrderDto
	err := ctx.ShouldBindJSON(&OrderDto)
	if err != nil {
		responses.ResponseClientError(ctx, err)
		return
	}
	orderId, err := h.service.Add(OrderDto)
	if err != nil {
		responses.ResponseClientError(ctx, err)
		return
	}
	responses.ResponseOkCreated(ctx, orderId)
}
func (h *orderHandler) getOneOrderById(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := h.service.ListOneById(id)
	if err != nil {
		responses.ResponseNotFound(ctx)
		return
	}
	responses.ResponseOkData(ctx, order)
}
