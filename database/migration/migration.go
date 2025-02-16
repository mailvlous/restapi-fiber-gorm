package migration

import (
	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/entity"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Users{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Migration successful")
	
}