package repositories

import (
	"preschool/models"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FindClasses() ([]models.Class, error)
	GetClass(ID int) (models.Class, error)
	GetClassUpdate(ID int) (models.Class, error)
	CreateClass(class models.Class) (models.Class, error)
	UpdateClass(class models.Class) (models.Class, error)
	DeleteClass(class models.Class) (models.Class, error)
}

func RepositoryClass(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindClasses() ([]models.Class, error) {
	var classes []models.Class
	err := r.db.Preload("Teacher").Preload("Enrollment").Find(&classes).Error

	return classes, err
}

func (r *repository) GetClass(ID int) (models.Class, error) {
	var class models.Class
	err := r.db.Preload("Teacher").Preload("Enrollment").First(&class, ID).Error

	return class, err
}

func (r *repository) GetClassUpdate(ID int) (models.Class, error) {
	var class models.Class
	err := r.db.First(&class, ID).Error

	return class, err
}

func (r *repository) CreateClass(class models.Class) (models.Class, error) {
	err := r.db.Create(&class).Error

	return class, err
}

func (r *repository) UpdateClass(class models.Class) (models.Class, error) {
	err := r.db.Save(&class).Error

	return class, err
}

func (r *repository) DeleteClass(class models.Class) (models.Class, error) {
	err := r.db.Delete(&class).Error

	return class, err
}
