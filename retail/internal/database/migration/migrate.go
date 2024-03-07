package migration

import (
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
