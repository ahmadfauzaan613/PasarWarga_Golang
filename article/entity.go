package article

import (
	"pasarwarga/category"
	"time"
)

type Article struct {
	ID         int
	Title      string
	Slug       string
	CategoryId int
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categori   category.Categori `gorm:"foreignKey:CategoryId"`
}
