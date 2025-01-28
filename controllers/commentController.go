package controllers

import (
	"net/http"
	initializers "project/initializer"
	"project/models"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var cmt models.Comment
	if err := c.ShouldBindJSON(&cmt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)
	userID := u.ID

	blog, _ := c.Get("blog")
	b := blog.(models.Blog)
	blogID := b.ID

	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cmt.UserID = userID 
	cmt.BlogID = blogID 

	if err := tx.Create(&cmt).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, cmt)

}

func UpdateComment(c *gin.Context) {
	id := c.Param("id")

	var userComment models.Comment
	user, _ := c.Get("user")
	u := user.(models.User)
	userID := u.ID

	if err := initializers.DB.Where("id = ?", userID).Find(&userComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog"})
		return
	}

	if userID != userComment.UserID{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})
		return
	}

	var cmt models.Comment
	if err := initializers.DB.First(&cmt, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := c.ShouldBindJSON(&cmt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&cmt)
	c.JSON(http.StatusOK, cmt)
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	var userComment models.Comment
	user, _ := c.Get("user")
	u := user.(models.User)
	userID := u.ID

	if err := initializers.DB.Where("id = ?", userID).Find(&userComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog"})
		return
	}

	if userID != userComment.UserID{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})
		return
	}

	//var cmt models.Comment

	if err := initializers.DB.Delete(&models.Comment{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
