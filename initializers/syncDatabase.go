package initializers

import (
	"fmt"
	"github.com/1mt142/verifier/models"
)

func SyncDatabase() {
	var err error
	err = DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		panic(err)
	}
	var tables []string
	result := DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Tables in database: %v\n", tables)
}
