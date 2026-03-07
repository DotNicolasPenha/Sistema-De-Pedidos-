package sistemadepedidos

import (
	"context"
	"time"

	"com.DotNicolasPenha.SistemaDePedidos/internal/repository/memory"
	"com.DotNicolasPenha.SistemaDePedidos/internal/service"
)

func main() {
	contextServices, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	repositoryOrder, repositoryPayment, repositoryProduct := memory.NewMemoryRepositorys()

	serviceOrder, servicePayment, serviceProduct := service.NewServices(
		contextServices, repositoryOrder, repositoryPayment, repositoryProduct,
	)
	println(serviceOrder.ListAll())
	println(servicePayment.ListAll())
	println(serviceProduct.ListAll())
}
