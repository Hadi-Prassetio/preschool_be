package models

import "time"
type Teacher struct {
	ID        int          `json:"id"  gorm:"primary_key:auto_increment"`
	FullName  string       `json:"fullname" gorm:"type: varchar(100)"`
	Email     string       `json:"email" gorm:"type: varchar(100)"`
	Phone     string       `json:"phone" gorm:"type: varchar(15)"`
	ClassID   int          `json:"class_id"`
	Class     ClassDetail  `json:"class"`
	AdminID   int          `json:"admin_id" gorm:"foreignKey:TeacherID"`
	Admin     AdminProfile `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}


type TeacherProfile struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (TeacherProfile) TableName() string {
	return "teachers"
}
type CreateTeacher struct {
	FullName string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	ClassID  int    `json:"class_id" validate:"required"`
	AdminID  int    `json:"admin_id"`
}

type UpdateTeacher struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	ClassID  int    `json:"class_id"`
	AdminID  int    `json:"admin_id"`
}