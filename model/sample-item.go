package model

type SampleItem struct {
	ID           uint      `gorm:"primary_key"`
	SampleItemID string    `gorm:"unique;not null"`
	Segments     []Segment `gorm:"foreignKey:MappingID"`
	ItemID       string
}

type Segment struct {
	ID        uint64 `gorm:"primaryKey"`
	Segment   string
	MappingID uint64
}
