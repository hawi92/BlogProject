package controllers

import (
	"net/http"
	initializers "project/initializer"
	"project/models"

	"github.com/gin-gonic/gin"
)

func CreateBlogRating(c *gin.Context) {
	var blog models.BlogRating
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)
	userID := u.ID

	blg, _ := c.Get("blog")
	b := blg.(models.Blog)
	blogID := b.ID

	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	blog.UserID = userID
	blog.BlogID = blogID 

	if err := tx.Create(&blog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create rating"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, blog)

}

