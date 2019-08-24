package models

import "time"

type DevPair struct {
	ID        uint `gorm:"primary_key" json:"id"`
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time `sql:"index"`
	Dev1      string
	Dev2      string
}
