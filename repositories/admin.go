package repositories

import (
	"preschool/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Register(admin models.Admin) (models.Admin, error)
	Login(email string) (models.Admin, error)
	Getadmin(ID int) (models.Admin, error)
}

func RepositoryAdmin(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(admin models.Admin) (models.Admin, error) {
	err := r.db.Create(&admin).Error
	return admin, err
}

func (r *repository) Login(email string) (models.Admin, error) {
	var admin models.Admin
	err := r.db.First(&admin, "email=?", email).Error

	return admin, err
}

func (r *repository) Getadmin(ID int) (models.Admin, error) {
	var admin models.Admin
	err := r.db.First(&admin, ID).Error

	return admin, err
}
