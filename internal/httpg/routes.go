package httpg

import (
	"com.DotNicolasPenha.SistemaDePedidos/internal/httpg/handlers"
	"com.DotNicolasPenha.SistemaDePedidos/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	g *gin.Engine,
	orderService service.OrderService,
	paymentService service.PaymentService,
	productService service.ProductService,
) {
	handlers.NewOrderHandler(g, orderService)
	handlers.NewProductHandler(g, productService)
	handlers.NewPaymentHandler(g, paymentService)
}
