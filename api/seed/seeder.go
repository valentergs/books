package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/valentergs/go-boilerplate/api/models"
)

var users = []models.User{
	models.User{
		Email:    "rodrigovalente@hotmail.com",
		Password: "Gustavo2012",
	},
	models.User{
		Email:    "vanessa@gmail.com",
		Password: "Gustavo2012",
	},
	models.User{
		Email:    "eduardo@gmail.com",
		Password: "Gustavo2012",
	},
	models.User{
		Email:    "gustavo@gmail.com",
		Password: "Gustavo2012",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		// posts[i].AuthorID = users[i].ID
	}

	// 	err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed posts table: %v", err)
	// 	}
	// }
}
