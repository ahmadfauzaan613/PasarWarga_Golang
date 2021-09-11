package category

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(category Categori) (Categori, error)
	FindAll() ([]Categori, error)
	FindByID(ID int) (Categori, error)
	UpdateCategory(category Categori) (Categori, error)
	CategoryDelete(category Categori) (Categori, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Categori) (Categori, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindAll() ([]Categori, error) {
	var categoris []Categori

	err := r.db.Find(&categoris).Error
	if err != nil {
		return categoris, err
	}

	return categoris, nil
}

func (r *repository) FindByID(ID int) (Categori, error) {
	var categoris Categori

	err := r.db.Where("id = ?", ID).Find(&categoris).Error
	if err != nil {
		return categoris, err
	}
	return categoris, nil
}

func (r *repository) UpdateCategory(category Categori) (Categori, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) CategoryDelete(category Categori) (Categori, error) {
	err := r.db.Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
