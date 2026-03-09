package main

import (
	"context"
	"time"

	"com.DotNicolasPenha.SistemaDePedidos/internal/httpg"
	"com.DotNicolasPenha.SistemaDePedidos/internal/repository/memory"
	"com.DotNicolasPenha.SistemaDePedidos/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// configure domain
	contextServices, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	repositoryOrder, repositoryPayment, repositoryProduct := memory.NewMemoryRepositorys()

	serviceOrder, servicePayment, serviceProduct := service.NewServices(
		contextServices, repositoryOrder, repositoryPayment, repositoryProduct,
	)
	// configure api
	g := gin.Default()
	httpg.SetupRoutes(g, serviceOrder, servicePayment, serviceProduct)
	g.Run(":8000")
}
