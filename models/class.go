package models

import "time"
type Class struct {
	ID         int          `json:"id"  gorm:"primary_key:auto_increment"`
	Name       string       `json:"name" gorm:"type: varchar(100)"`
	Capacity   int          `json:"capacity" gorm:"type: int(2)"`
	Teacher    []Teacher    `json:"teacher" gorm:"foreignKey:ClassID"`
	Enrollment []Enrollment `json:"enroll" gorm:"foreignKey:ClassID"`
	AdminID    int          `json:"admin_id" gorm:"foreignKey:ClassID"`
	Admin      AdminProfile `json:"-" gorm:"foreignKey:AdminID"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
}

type ClassDetail struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

func (ClassDetail) TableName() string {
	return "classes"
}
type CreateClass struct {
	Name     string `json:"name" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
}
type UpdateClass struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	AdminID  int    `json:"admin_id"`
}