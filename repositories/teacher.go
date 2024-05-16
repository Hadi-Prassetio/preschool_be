package repositories

import (
	"preschool/models"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	FindTeachers() ([]models.Teacher, error)
	GetTeacher(ID int) (models.Teacher, error)
	GetTeacherUpdate(ID int) (models.Teacher, error)
	CreateTeacher(teacher models.Teacher) (models.Teacher, error)
	UpdateTeacher(teacher models.Teacher) (models.Teacher, error)
	DeleteTeacher(teacher models.Teacher) (models.Teacher, error)
}

func RepositoryTeacher(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTeachers() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := r.db.Preload("Class").Find(&teachers).Error

	return teachers, err
}

func (r *repository) GetTeacher(ID int) (models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.Preload("Class").First(&teacher, ID).Error

	return teacher, err
}

func (r *repository) GetTeacherUpdate(ID int) (models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.First(&teacher, ID).Error

	return teacher, err
}


func (r *repository) CreateTeacher(teacher models.Teacher) (models.Teacher, error) {
	err := r.db.Create(&teacher).Error

	return teacher, err
}

func (r *repository) UpdateTeacher(teacher models.Teacher) (models.Teacher, error) {
	err := r.db.Save(&teacher).Error

	return teacher, err
}

func (r *repository) DeleteTeacher(teacher models.Teacher) (models.Teacher, error) {
	err := r.db.Delete(&teacher).Error

	return teacher, err
}
