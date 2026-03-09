package handlers

import (
	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"com.DotNicolasPenha.SistemaDePedidos/internal/httpg/responses"
	"com.DotNicolasPenha.SistemaDePedidos/internal/service"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(g *gin.Engine, productService service.ProductService) {
	productsRoute := g.Group("/products")
	productHandler := productHandler{service: productService}
	{
		productsRoute.GET("/", productHandler.getProducts)
		productsRoute.GET("/:id", productHandler.getOneProductById)
		productsRoute.POST("/", productHandler.postProduct)
	}
}

func (h *productHandler) getProducts(ctx *gin.Context) {
	products, err := h.service.ListAll()
	if err != nil {
		responses.ReponseInternalError(ctx, err)
		return
	}
	responses.ResponseOkData(ctx, products)
}
func (h *productHandler) postProduct(ctx *gin.Context) {
	var ProductDto domain.CreateProductDto
	err := ctx.ShouldBindJSON(&ProductDto)
	if err != nil {
		responses.ResponseClientError(ctx, err)
		return
	}
	productId, err := h.service.Add(ProductDto)
	if err != nil {
		responses.ResponseClientError(ctx, err)
		return
	}
	responses.ResponseOkCreated(ctx, productId)
}
func (h *productHandler) getOneProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.service.ListOneById(id)
	if err != nil {
		responses.ResponseNotFound(ctx)
		return
	}
	responses.ResponseOkData(ctx, product)
}
