package model

import "time"

type Hello struct {
	ID int
	Name string
	Age int

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
