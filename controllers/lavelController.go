package controllers

import (
	"fmt"
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LavelCreate(c *gin.Context) {

	// self-referencing hierarchical table using GORM:
	
	// Create new root category
	//root := &models.TypeTree{Name: "Fashion"}
	//result := initializers.DB.Create(&root)
	//if result.Error != nil {
	//	panic("failed to create root category")
	//}
	//fmt.Println("Created Root Category")

	// make a pointer value
	var ptr *uint
	val := uint(3)
	ptr = &val
	fmt.Println("ptr val:", ptr)

	// Create new child category
	child := &models.TypeTree{Name: "Unisex", ParentID: ptr}
	result := initializers.DB.Create(&child)
	if result.Error != nil {
		panic("failed to create child category")
	}
	fmt.Println("Created Child Category")

	// Retrieve all categories with their parents and children
	var categories []*models.TypeTree
	result = initializers.DB.Preload("Parent").Preload("Children").Find(&categories)
	if result.Error != nil {
		panic("failed to retrieve categories")
	}
	fmt.Println("Retrieved Categories")
	fmt.Println(categories)

	c.JSON(http.StatusOK, gin.H{
		"message": categories,
	})

}
