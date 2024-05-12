package database

import (
	"fmt"
	"preschool/models"
	"preschool/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Admin{}, &models.Teacher{})

	if err != nil {
		fmt.Println(err)
		panic("migration error")
	}
	fmt.Println("migration success")
}
