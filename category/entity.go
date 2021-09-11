package category

import "time"

type Categori struct {
	ID           int
	CategoryName string
	CategorySlug string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
