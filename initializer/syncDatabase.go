package initializers

import "project/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Blog{}, &models.BlogRating{},&models.Comment{}, &models.Like{})
}
