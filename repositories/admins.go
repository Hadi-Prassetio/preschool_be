package repositories

import (
	"preschool/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAdmins() ([]models.Admin, error)
	GetAdmin(ID int) (models.Admin, error)
	UpdateAdmin(admin models.Admin) (models.Admin, error)
}

func RepositoryAdmin(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAdmins() ([]models.Admin, error) {
	var admins []models.Admin
	err := r.db.Find(&admins).Error

	return admins, err
}

func (r *repository) GetAdmin(ID int) (models.Admin, error) {
	var admin models.Admin
	err := r.db.First(&admin, ID).Error

	return admin, err
}

func (r *repository) UpdateAdmin(admin models.Admin) (models.Admin, error) {
	err := r.db.Save(&admin).Error

	return admin, err
}
