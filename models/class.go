package models

import "time"

type Class struct {
	ID         int          `json:"id"  gorm:"primary_key:auto_increment"`
	Name       string       `json:"name" gorm:"type: varchar(100)"`
	Capacity   int          `json:"capacity" gorm:"type: int(2)"`
	Teacher    []Teacher    `json:"teacher"`
	Enrollment []Enrollment `json:"enroll"`
	AdminID    int          `json:"admin_id"`
	Admin      AdminProfile `json:"admin"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
}

type ClassDetail struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Capacity  string		 `json:"capacity"`
}

func (ClassDetail) TableName() string {
	return "classes"
}
