package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	Title      string `json:"title"`
	Content    string  `json:"content"`
}

type BlogRating struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	BlogID      uint `json:"blog_id"`
	RatingValue    int  `json:"rating_value"`
}
type Comment struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	BlogID      uint `json:"blog_id"`
	Content    int  `json:"content"`
}
type Like struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	BlogID      uint `json:"blog_id"`
}





