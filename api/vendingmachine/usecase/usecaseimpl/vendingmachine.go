package usecaseimpl

import (
	"context"
	"fmt"
	"perqara_testing/api/vendingmachine/repository"
	"perqara_testing/api/vendingmachine/usecase"
	"perqara_testing/ent"
	"strconv"
	"time"

	"github.com/uptrace/bun/schema"
)

type VendingMachineService struct {
	vendingMachineRepo repository.VendingMachineRepository
}

func NewVendingMachineService(vendingMachineRepo repository.VendingMachineRepository) *VendingMachineService {
	return &VendingMachineService{
		vendingMachineRepo: vendingMachineRepo,
	}
}

func (s *VendingMachineService) ItemAll(ctx context.Context) (dto []*usecase.VendingMachine, err error) {
	items, itemsErr := s.vendingMachineRepo.ItemList(ctx)
	if itemsErr != nil {
		err = itemsErr
		return
	}

	dto = make([]*usecase.VendingMachine, len(items))
	for i, value := range items {
		dto[i] = &usecase.VendingMachine{
			ID:    fmt.Sprintf("%d", value.ID),
			Item:  value.Item,
			Price: fmt.Sprintf("%d", value.Price),
		}
	}

	return
}

func (s *VendingMachineService) CreateItem(ctx context.Context, cmd *usecase.VendingMachineRequestCommand) (err error) {
	if err = itemCreateValidation(*cmd); err != nil {
		return
	}

	price, _ := strconv.ParseInt(cmd.Price, 10, 64)

	ent := ent.VendingMachines{
		Item:  cmd.Item,
		Price: price,
	}

	if err = s.vendingMachineRepo.CreateItem(ctx, &ent); err != nil {
		return
	}

	return
}

func (s *VendingMachineService) DetailItem(ctx context.Context, itemId string) (dto *usecase.VendingMachine, err error) {
	idItem, itemIdErr := strconv.ParseInt(itemId, 10, 64)
	if itemIdErr != nil {
		err = itemIdErr
		return
	}

	ent := ent.VendingMachines{
		ID: idItem,
	}

	dta, dtaErr := s.vendingMachineRepo.DetailItem(ctx, &ent)
	if dtaErr != nil {
		err = dtaErr
		return
	}

	dto = &usecase.VendingMachine{
		ID:    fmt.Sprintf("%d", dta.ID),
		Item:  dta.Item,
		Price: fmt.Sprintf("%d", dta.Price),
	}

	return
}

func (s *VendingMachineService) UpdateItem(ctx context.Context, cmd *usecase.VendingMachineUpdateCommand) (err error) {
	if err = itemUpdateValidation(*cmd); err != nil {
		return
	}

	price, _ := strconv.ParseInt(cmd.Price, 10, 64)
	id, _ := strconv.ParseInt(cmd.ID, 10, 64)

	ent := ent.VendingMachines{ID: id}

	findVandingMachine, findVandingMachineErr := s.vendingMachineRepo.DetailItem(ctx, &ent)
	if findVandingMachineErr != nil {
		err = findVandingMachineErr
		return
	}

	ent.Item = cmd.Item
	ent.Price = price
	ent.CreatedAt = findVandingMachine.CreatedAt
	ent.UpdatedAt = schema.NullTime{
		Time: time.Now().Local(),
	}

	if err = s.vendingMachineRepo.UpdateItem(ctx, &ent); err != nil {
		return
	}

	return
}

func (s *VendingMachineService) DeleteItem(ctx context.Context, itemId string) (err error) {
	id, _ := strconv.ParseInt(itemId, 10, 64)

	ent := ent.VendingMachines{
		ID: id,
	}

	if _, err = s.vendingMachineRepo.DetailItem(ctx, &ent); err != nil {
		return
	}

	if err = s.vendingMachineRepo.DeleteItem(ctx, itemId); err != nil {
		return
	}

	return
}
