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
	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", middleware.RequireAuth, controllers.GetAllPosts)
	r.GET("/post/:id", middleware.RequireAuth, controllers.GetPost)
	r.PUT("/post/:id", middleware.RequireAuth, controllers.UpdatePost)
	r.DELETE("/post/:id", middleware.RequireAuth, controllers.DeletePost)
	// Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("otp-verify", controllers.OtpVerify)
	// User
	r.GET("/users", middleware.RequireAuth, controllers.GetUsers)
	// Random
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
