package main

import (
	"github.com/1mt142/verifier/models"
	"net/http"
	"os"

	"github.com/1mt142/verifier/controllers"
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	// initializers.ConnectToRedis()
	//initializers.SyncDatabase()
	// Set up Zerolog logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	//

}

// Global Logger

func main() {
	r := gin.Default()

	// logger

	//Get all env variables
	//fmt.Println(os.Environ())

	// Post
	r.POST("api/v1/post", controllers.PostCreate)
	r.GET("api/v1/posts", middleware.RequireAuth, controllers.GetAllPosts)
	r.GET("api/v1/post/:id", middleware.RequireAuth, controllers.GetPost)
	r.PUT("api/v1/post/:id", middleware.RequireAuth, controllers.UpdatePost)
	r.DELETE("api/v1/post/:id", middleware.RequireAuth, controllers.DeletePost)
	// Auth
	r.POST("api/v1/signup", controllers.Signup)
	r.POST("api/v1/login", controllers.Login)
	r.POST("api/v1/otp-verify", controllers.OtpVerify)
	// User
	r.GET("api/v1/users", middleware.RequireAuth, controllers.GetUsers)
	// Random
	r.GET("api/v1/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("api/v1/relation", controllers.RelationTest)
	r.POST("api/v1/article", controllers.CreateArticle)
	r.GET("api/v1/article/:id", controllers.FetchArticle)
	r.GET("api/v1/article/tag/:id", controllers.FetchArticleByTag)
	r.GET("api/v1/article/category/:id", controllers.FetchArticleByCategory)
	r.GET("api/v1/test", func(c *gin.Context) {

		// var user *models.User
		// var service *models.Service

		// err := initializers.DB.Model(&models.User{}).Where("id = ?", "3838a4ce-374a-427d-935b-75f757a4e28b").Association("Services").Append(&models.Service{
		// 	Name:        "test",
		// 	Address:     "test Addr",
		// 	CompanyType: "string",
		// }).Error()

		// fmt.Println(err)

		// fmt.Printf("%#v", user)

		// ----------- if i want to

		// create a new article
		//newArticle := models.Article{Title: "New Article", Content: "This is the content of the new article."}
		//// associate the article with a category
		//category := models.Category{Name: "Science"}
		//initializers.DB.Create(&category)
		//newArticle.CategoryID = category.ID
		//// associate the article with some tags
		//tags := []models.Tag{{Name: "Golang"}, {Name: "Database"}}
		//initializers.DB.Create(&tags)
		//for _, tag := range tags {
		//	newArticle.Tags = append(newArticle.Tags, &tag)
		//}
		//// save the article to the database
		//initializers.DB.Create(&newArticle)

		// if i have data

		// retrieve an existing category from the database
		var category models.Category
		initializers.DB.First(&category, "name = ?", "Science")

		// retrieve existing tags from the database
		var tags []models.Tag
		initializers.DB.Where("name IN (?)", []string{"Golang", "Database"}).Find(&tags)

		println("Tags", tags)

		// create a new article and associate it with the existing category and tags
		newArticle := models.Article{
			Title:      "This is My First Relationship article in golang app",
			Content:    "I want to write so many things here but the problem is ,I dont know how to make  relation among db",
			CategoryID: category.ID,
			Tags: []*models.Tag{
				&tags[0],
				&tags[1],
			}}

		// save the new article to the database
		initializers.DB.Create(&newArticle)

		c.JSON(http.StatusOK, gin.H{
			"message": "You are logged in!",
			"data":    newArticle,
		})

	})

	r.Run()
}
