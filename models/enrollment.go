package models

import "time"

type Enrollment struct {
	ID         int          `json:"id"  gorm:"primary_key:auto_increment"`
	FatherName string       `json:"father_name" gorm:"type: varchar(100)"`
	MotherName string       `json:"mother_name" gorm:"type: varchar(100)"`
	Email      string       `json:"email" gorm:"type: varchar(100)"`
	Phone      string       `json:"phone" gorm:"type: varchar(15)"`
	ChildName  string       `json:"child_name" gorm:"type: varchar(100)"`
	ChildAge   int          `json:"child_age" gorm:"type: int(2)"`
	Status     string       `json:"status" gorm:"type:varchar(100)"`
	ClassID    int          `json:"class_id"`
	Class      ClassDetail  `json:"class"`
	AdminID    int          `json:"admin_id"`
	Admin      AdminProfile `json:"admin"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
}

type EnrollmentDetail struct {
	ID         int    `json:"id"`
	FatherName string `json:"father_name"`
	MotherName string `json:"mother_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	ChildName  string `json:"child_name"`
	ChildAge   int    `json:"child_age"`
	Status     string `json:"status"`
}

func (EnrollmentDetail) TableName() string {
	return "enrollments"
}
