package seed

import (
	"log"

	"github.com/bmd-rezafahlevi/fullstack/api/models"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Nickname: "Steven Victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	{
		Title:   "Title 1",
		Content: "Hello World 1",
	},
	{
		Title:   "Title 2",
		Content: "Hello Wolrd 2",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
