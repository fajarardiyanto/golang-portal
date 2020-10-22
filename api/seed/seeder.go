package seed

import (
	"log"
	"rest-api-tutorial/portal/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Username: "Fajar Ardiyanto",
		Email:    "fajarardiyanto.web@gmail.com",
		Password: "fajar123",
	},
	models.User{
		Username: "Ryuusei",
		Email:    "pythonersdjango@gmail.com",
		Password: "ryuusei123",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello World",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello World 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
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
