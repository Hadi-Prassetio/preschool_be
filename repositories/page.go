package repositories

import (
	"preschool/models"

	"gorm.io/gorm"
)

type PageRepository interface {
	FindPages() ([]models.Page, error)
	GetPage(ID int) (models.Page, error)
	CreatePage(page models.Page) (models.Page, error)
	UpdatePage(page models.Page) (models.Page, error)
	DeletePage(page models.Page) (models.Page, error)
}

func RepositoryPage(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPages() ([]models.Page, error) {
	var pages []models.Page
	err := r.db.Find(&pages).Error

	return pages, err
}

func (r *repository) GetPage(ID int) (models.Page, error) {
	var page models.Page
	err := r.db.First(&page, ID).Error

	return page, err
}

func (r *repository) CreatePage(page models.Page) (models.Page, error) {
	err := r.db.Create(&page).Error

	return page, err
}

func (r *repository) UpdatePage(page models.Page) (models.Page, error) {
	err := r.db.Save(&page).Error

	return page, err
}

func (r *repository) DeletePage(page models.Page) (models.Page, error) {
	err := r.db.Delete(&page).Error

	return page, err
}
