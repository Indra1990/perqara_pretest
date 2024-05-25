package repositorybun

import (
	"context"
	"perqara_testing/ent"

	"github.com/uptrace/bun"
)

type VendingMachineRepository struct {
	db *bun.DB
}

func NewVendingMachineRepository(db *bun.DB) *VendingMachineRepository {
	return &VendingMachineRepository{
		db: db,
	}
}

func (r *VendingMachineRepository) ItemList(ctx context.Context) ([]*ent.VendingMachines, error) {
	ents := []*ent.VendingMachines{}

	if err := r.db.NewSelect().Model(&ents).Scan(ctx); err != nil {
		return nil, err
	}

	return ents, nil
}

func (r *VendingMachineRepository) CreateItem(ctx context.Context, ent *ent.VendingMachines) error {
	if _, err := r.db.NewInsert().Model(ent).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *VendingMachineRepository) DetailItem(ctx context.Context, ent *ent.VendingMachines) (*ent.VendingMachines, error) {

	if err := r.db.NewSelect().Model(ent).WherePK().Scan(ctx); err != nil {
		return nil, err
	}

	return ent, nil
}

func (r *VendingMachineRepository) UpdateItem(ctx context.Context, ent *ent.VendingMachines) error {
	_, err := r.db.NewUpdate().Model(ent).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *VendingMachineRepository) DeleteItem(ctx context.Context, itemId string) error {
	ent := &ent.VendingMachines{}
	_, err := r.db.NewDelete().Model(ent).Where("id = ?", itemId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
