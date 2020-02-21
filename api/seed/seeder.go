package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/valentergs/booksv2/api/models"
)

var users = []models.User{
	models.User{
		User:     "rova1976",
		Email:    "valentergs@gmail.com",
		Password: "Gustavo2012",
		Admin:    true,
	},
	models.User{
		User:     "vany1981",
		Email:    "vanessa@gmail.com",
		Password: "Gustavo2012",
		Admin:    false,
	},
}

var books = []models.Book{
	models.Book{
		ISBN:   "9788554126254",
		Titulo: "O Sol já Brilhou!!!",
		Autor:  "Anthony Ray Hinton",
		Slug:   "o-sol-ainda-brilha",
		Cdd:    "823",
		Capa:   "https://images-na.ssl-images-amazon.com/images/I/81BcnC+VFYL.jpg",
	},
	models.Book{
		ISBN:   "9788551005767",
		Titulo: "A Quietude é a Chave",
		Autor:  "Ryan Holiday",
		Slug:   "a-quietude-é-a-chave",
		Cdd:    "500",
		Capa:   "https://images-na.ssl-images-amazon.com/images/I/81BcnC+VFYL.jpg",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Book{}).Error
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
	}

	for i, _ := range books {
		err = db.Debug().Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}
	}
}
