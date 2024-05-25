package ent

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type VendingMachines struct {
	bun.BaseModel `bun:"table:vending_machines"`
	ID            int64     `bun:"id,pk,autoincrement"`
	Item          string    `bun:"item"`
	Price         int64     `bun:"price"`
	CreatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt     bun.NullTime
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}

var _ bun.AfterCreateTableHook = (*VendingMachines)(nil)

func (vendingMachine *VendingMachines) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	// insertData := []map[string]interface{}{
	// 	{
	// 		"item":  "hjghfgh",
	// 		"price": 909,
	// 	},
	// 	{
	// 		"item":  "fhgfhgf",
	// 		"price": 909,
	// 	},
	// 	{
	// 		"item":  "fghgfh",
	// 		"price": 909,
	// 	},
	// 	{
	// 		"item":  "fghfhfgh",
	// 		"price": 909,
	// 	},
	// 	{
	// 		"item":  "gfgfhfdhfd",
	// 		"price": 909,
	// 	},
	// }

	sdfdsf := []VendingMachines{
		VendingMachines{
			Item:  "dsfdsf",
			Price: 45454,
		},
	}

	_, err := query.NewInsert().Model(sdfdsf).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
