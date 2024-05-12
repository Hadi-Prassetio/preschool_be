package models

import "time"

type Teacher struct {
	ID        int       `json:"id"  gorm:"primary_key:auto_increment"`
	FullName  string    `json:"fullname" gorm:"type: varchar(100)"`
	Email     string    `json:"email" gorm:"type: varchar(70)"`
	Phone     string    `json:"phone" gorm:"type: varchar(15)"`
	Subject   string    `json:"subject" gorm:"type: varchar(100)"`
	Register  string    `json:"register" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TeacherAdmin struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Subject  string `json:"subject"`
}

func (TeacherAdmin) TableName() string {
	return "teachers"
}
