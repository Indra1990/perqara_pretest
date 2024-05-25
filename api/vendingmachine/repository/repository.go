package repository

import (
	"context"
	"perqara_testing/ent"
)

type VendingMachineRepository interface {
	ItemList(ctx context.Context) ([]*ent.VendingMachines, error)
	CreateItem(ctx context.Context, ent *ent.VendingMachines) error
	DetailItem(ctx context.Context, ent *ent.VendingMachines) (*ent.VendingMachines, error)
	UpdateItem(ctx context.Context, ent *ent.VendingMachines) error
	DeleteItem(ctx context.Context, itemId string) error
}
