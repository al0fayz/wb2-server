package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"wb2/api/models"
)

var roles = []models.Role{
	models.Role{
		Name: "Admin",
	},
	models.Role{
		Name: "Mitra",
	},
	models.Role{
		Name: "Agent",
	},
	models.Role{
		Name: "Sub Agent",
	},
}
var users = []models.User{
	models.User{
		Username: "admin",
		Email:    "admin@me.com",
		Password: "Pandi123#",
	},
}
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Role{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Role{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range roles {
		err = db.Debug().Model(&models.Role{}).Create(&roles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed roles table: %v", err)
		}
	}

	for i, _ := range users {
		users[i].RoleID = 1
	
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}