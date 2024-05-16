package repositories

import (
	"preschool/models"

	"gorm.io/gorm"
)

type VisitorRepository interface {
	FindVisitors() ([]models.Visitor, error)
	FindVisitorStatus(Status string) ([]models.Visitor, error)
	GetVisitor(ID int) (models.Visitor, error)
	CreateVisitor(visitor models.Visitor) (models.Visitor, error)
	UpdateVisitor(visitor models.Visitor) (models.Visitor, error)
	DeleteVisitor(visitor models.Visitor) (models.Visitor, error)
}

func RepositoryVisitor(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindVisitors() ([]models.Visitor, error) {
	var visitors []models.Visitor
	err := r.db.Find(&visitors).Error

	return visitors, err
}

func (r *repository) FindVisitorStatus(Status string) ([]models.Visitor, error) {
	var visitors []models.Visitor
	err := r.db.Find(&visitors, "status = ?", Status).Error

	return visitors, err
}

func (r *repository) GetVisitor(ID int) (models.Visitor, error) {
	var visitor models.Visitor
	err := r.db.First(&visitor, ID).Error

	return visitor, err
}

func (r *repository) CreateVisitor(visitor models.Visitor) (models.Visitor, error) {
	err := r.db.Create(&visitor).Error

	return visitor, err
}

func (r *repository) UpdateVisitor(visitor models.Visitor) (models.Visitor, error) {
	err := r.db.Save(&visitor).Error

	return visitor, err
}

func (r *repository) DeleteVisitor(visitor models.Visitor) (models.Visitor, error) {
	err := r.db.Delete(&visitor).Error

	return visitor, err
}
