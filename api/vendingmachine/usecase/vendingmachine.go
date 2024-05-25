package usecase

import (
	"context"
)

type (
	VendingMachineRequestCommand struct {
		Item  string `json:"item"`
		Price string `json:"price"`
	}

	VendingMachineUpdateCommand struct {
		ID    string `json:"id,omitempty" swaggerignore:"true"`
		Item  string `json:"item"`
		Price string `json:"price"`
	}
)

type VendingMachineService interface {
	ItemAll(ctx context.Context) (dto []*VendingMachine, err error)
	CreateItem(ctx context.Context, cmd *VendingMachineRequestCommand) (err error)
	DetailItem(ctx context.Context, itemId string) (dto *VendingMachine, err error)
	UpdateItem(ctx context.Context, cmd *VendingMachineUpdateCommand) (err error)
	DeleteItem(ctx context.Context, itemId string) (err error)
}
