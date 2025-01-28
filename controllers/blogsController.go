package controllers

import (
	"net/http"
	initializers "project/initializer"
	"project/models"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)
	userID := u.ID

	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	blog.UserID = userID 

	if err := tx.Create(&blog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, blog)

}

func UpdateBlog(c *gin.Context) {
	id := c.Param("id")

	var blogUser models.Blog
	user, _ := c.Get("user")
	u := user.(models.User)
	userID := u.ID

	if err := initializers.DB.Where("user_id = ?", userID).Find(&blogUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog"})
		return
	}

	if userID != blogUser.UserID{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})
		return
	}

	var blog models.Blog
	if err := initializers.DB.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&blog)
	c.JSON(http.StatusOK, blog)
}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	//var blog models.Blog

	if err := initializers.DB.Delete(&models.Blog{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetBlogs(c *gin.Context) {
	var blogs []models.Blog

	user, _ := c.Get("user")
	u := user.(models.User)
	userID := u.ID

	if err := initializers.DB.Where("user_id = ?", userID).Find(&blogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blogs"})
		return
	}
	c.JSON(http.StatusOK, blogs)

}

func GetBlogsByID(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	if err := initializers.DB.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}


