package model

import "github.com/lib/pq"

type SampleItem struct {
	ID           uint64 `gorm:"primary_key"`
	SampleItemID string
	Segments     pq.StringArray `gorm:"type:text[]"`
	ItemID       string
}
