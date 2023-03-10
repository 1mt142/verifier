package main

import (
	"github.com/1mt142/verifier/controllers"
	"github.com/1mt142/verifier/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	//Get all env variables
	//fmt.Println(os.Environ())

	// post
	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	// user
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.Validate)

	r.Run()
}
