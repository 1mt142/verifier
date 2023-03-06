package controllers

import (
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
	"github.com/gin-gonic/gin"
)

var body struct {
	Title string
	Body  string
}

func PostCreate(c *gin.Context) {
	// get data
	c.Bind(&body)
	// create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(200)
		return
	}
	c.JSON(200, gin.H{
		"data": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"data": posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.Find(&post, id)
	c.JSON(200, gin.H{
		"data": post,
	})
}

func UpdatePost(c *gin.Context) {
	// get id
	id := c.Param("id")
	// get body
	c.Bind(&body)
	// find data
	var post models.Post
	initializers.DB.Find(&post, id)
	// update data
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	c.JSON(200, gin.H{
		"message": "Data updated successfully",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	// get id
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	c.JSON(200, gin.H{
		"message": "Data deleted successfully",
	})

}
