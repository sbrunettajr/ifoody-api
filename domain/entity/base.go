package entity

import "time"

type base struct {
	ID        uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
