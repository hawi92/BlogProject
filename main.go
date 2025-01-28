package main

import (
	"log"
	"net/http"
	"project/controllers"
	initializers "project/initializer"
	"project/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()

	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	blogRoutes := r.Group("/blog")
	blogRoutes.Use(middleware.RequiredAuth)
	{
		blogRoutes.POST("", controllers.CreateBlog)
		blogRoutes.GET("", controllers.GetBlogs)
		blogRoutes.GET("/:id", controllers.GetBlogsByID)
		blogRoutes.PUT("/:id", controllers.UpdateBlog)
		blogRoutes.DELETE("/:id", controllers.DeleteBlog)
	}

	commentRoutes := r.Group("/comment")
	commentRoutes.Use(middleware.RequiredAuth)
	{
		commentRoutes.POST("/:id", controllers.CreateComment)
		commentRoutes.PUT("/:id", controllers.UpdateComment)
		commentRoutes.DELETE("/:id", controllers.DeleteComment)
	}

	ratingRoutes := r.Group("/rate")
	ratingRoutes.Use(middleware.RequiredAuth)
	{
		ratingRoutes.POST("", controllers.CreateBlogRating)
	}

	r.Run()
	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}

}
