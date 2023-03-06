package main

import (
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
