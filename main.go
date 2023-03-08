package main

import (
	"fmt"
	"github.com/1mt142/verifier/controllers"
	"github.com/1mt142/verifier/initializers"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("It is ok.")
	fmt.Println("Hi")
	r := gin.Default()

	//Get all env variables
	fmt.Println(os.Environ())

	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.Run()
}
