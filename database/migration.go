package database

import (
	"fmt"
	"preschool/models"
	"preschool/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Admin{}, &models.Teacher{}, &models.Page{}, &models.Enrollment{}, &models.Class{}, &models.Visitor{})

	if err != nil {
		fmt.Println(err)
		panic("migration error")
	}
	fmt.Println("migration success")
}
