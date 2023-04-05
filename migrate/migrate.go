package main

import (
	"fmt"

	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	var err error

	err = initializers.DB.AutoMigrate(&models.TypeTree{}, &models.User{}, &models.Post{}, &models.OTP{}, &models.Service{}, &models.Article{}, &models.Category{}, &models.Tag{})

	//
	if err != nil {
		panic(err)
	}
	var tables []string
	result := initializers.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Tables in database: %v\n", tables)

}
