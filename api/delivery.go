package api

import (
	"perqara_testing/api/vendingmachine"
	"perqara_testing/api/vendingmachine/repository/repositorybun"
	"perqara_testing/api/vendingmachine/usecase/usecaseimpl"

	"github.com/uptrace/bun"
)

// @title			Test Perqara API
// @version			2.0
// @This			Test Perqara
// @termsOfService	http://swagger.io/terms/
// @contact.name	Muhammad Indra
// @contact.email	muhammad.indra65@gmail.com
// @license.name	Test Perqara API
// @license.url
// @host			localhost:8081
// @BasePath		/
func deliveryVendingMachine(db *bun.DB) *vendingmachine.VendingMachineHttpRouterRegistry {
	repoVendingMachine := repositorybun.NewVendingMachineRepository(db)
	serviceVendingMachine := usecaseimpl.NewVendingMachineService(repoVendingMachine)
	vendingMachineRoute := vendingmachine.NewTransactionHttpRouterRegistry(serviceVendingMachine)

	return vendingMachineRoute
}
func DeliveryTestVendingMachine(db *bun.DB) *vendingmachine.VendingMachineHttpRouterRegistry {
	repoVendingMachine := repositorybun.NewVendingMachineRepository(db)
	serviceVendingMachine := usecaseimpl.NewVendingMachineService(repoVendingMachine)
	vendingMachineRoute := vendingmachine.NewTransactionHttpRouterRegistry(serviceVendingMachine)

	return vendingMachineRoute
}
