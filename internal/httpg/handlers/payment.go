package handlers

import (
	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"com.DotNicolasPenha.SistemaDePedidos/internal/httpg/responses"
	"com.DotNicolasPenha.SistemaDePedidos/internal/service"
	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(g *gin.Engine, paymentService service.PaymentService) {
	paymentsRoute := g.Group("/payments")
	paymentHandler := paymentHandler{service: paymentService}
	{
		paymentsRoute.GET("/", paymentHandler.getPayments)
		paymentsRoute.GET("/:id", paymentHandler.getOnePaymentById)
		paymentsRoute.POST("/", paymentHandler.postPayment)
	}
}

func (h *paymentHandler) getPayments(ctx *gin.Context) {
	payments, err := h.service.ListAll()
	if err != nil {
		responses.ReponseInternalError(ctx, err)
		return
	}
	responses.ResponseOkData(ctx, payments)
}
func (h *paymentHandler) postPayment(ctx *gin.Context) {
	var PaymentDto domain.CreatePaymentDto
	err := ctx.ShouldBindJSON(&PaymentDto)
	if err != nil {
		responses.ResponseClientError(ctx, err)
		return
	}
	paymentId, err := h.service.Add(PaymentDto)
	if err != nil {
		responses.ResponseClientError(ctx, err)
		return
	}
	responses.ResponseOkCreated(ctx, paymentId)
}
func (h *paymentHandler) getOnePaymentById(ctx *gin.Context) {
	id := ctx.Param("id")
	payment, err := h.service.ListOneById(id)
	if err != nil {
		responses.ResponseNotFound(ctx)
		return
	}
	responses.ResponseOkData(ctx, payment)
}
