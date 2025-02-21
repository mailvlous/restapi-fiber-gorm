package migration

import (
	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/entity"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Users{}, &entity.Books{}, &entity.Photos{}, &entity.Categories{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Migration successful")
	
}