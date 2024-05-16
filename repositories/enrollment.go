package repositories

import (
	"preschool/models"

	"gorm.io/gorm"
)

type EnrollmentRepository interface {
	FindEnrollments() ([]models.Enrollment, error)
	FindEnrollmentStatus(Status string) ([]models.Enrollment, error)
	GetEnrollment(ID int) (models.Enrollment, error)
	GetEnrollmentUpdate(ID int) (models.Enrollment, error)
	CreateEnrollment(enrollment models.Enrollment) (models.Enrollment, error)
	UpdateEnrollment(enrollment models.Enrollment) (models.Enrollment, error)
	DeleteEnrollment(enrollment models.Enrollment) (models.Enrollment, error)
}

func RepositoryEnrollment(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEnrollments() ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := r.db.Preload("Class").Find(&enrollments).Error

	return enrollments, err
}

func (r *repository) FindEnrollmentStatus(Status string) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := r.db.Find(&enrollments, "status = ?", Status).Error

	return enrollments, err
}

func (r *repository) GetEnrollment(ID int) (models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.Preload("Class").First(&enrollment, ID).Error

	return enrollment, err
}

func (r *repository) GetEnrollmentUpdate(ID int) (models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.First(&enrollment, ID).Error

	return enrollment, err
}

func (r *repository) CreateEnrollment(enrollment models.Enrollment) (models.Enrollment, error) {
	err := r.db.Create(&enrollment).Error

	return enrollment, err
}

func (r *repository) UpdateEnrollment(enrollment models.Enrollment) (models.Enrollment, error) {
	err := r.db.Save(&enrollment).Error

	return enrollment, err
}

func (r *repository) DeleteEnrollment(enrollment models.Enrollment) (models.Enrollment, error) {
	err := r.db.Delete(&enrollment).Error

	return enrollment, err
}
