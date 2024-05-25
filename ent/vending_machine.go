package ent

import (
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
