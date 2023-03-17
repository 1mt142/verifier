package main

import (
	"github.com/1mt142/verifier/controllers"
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
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

	r.Run()
}
