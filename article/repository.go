package article

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(article Article) (Article, error)
	FindAll() ([]Article, error)
	FindByID(ID int) (Article, error)
	UpdateArticle(article Article) (Article, error)
	ArticleDelete(article Article) (Article, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(article Article) (Article, error) {
	err := r.db.Create(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) FindAll() ([]Article, error) {
	var articles []Article

	err := r.db.Preload("Categori").Find(&articles).Error
	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (r *repository) FindByID(ID int) (Article, error) {
	var article Article

	err := r.db.Preload("Categori").Where("id = ?", ID).Find(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}

func (r *repository) UpdateArticle(article Article) (Article, error) {
	err := r.db.Save(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) ArticleDelete(article Article) (Article, error) {
	err := r.db.Delete(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}
