package controllers

import (
	"fmt"
	"github.com/1mt142/verifier/initializers"
	"github.com/1mt142/verifier/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RelationTest(c *gin.Context) {
	var db = initializers.DB

	// retrieve an existing category from the database
	var category models.Category
	initializers.DB.First(&category, "name = ?", "Science")

	// retrieve existing tags from the database
	var tags []models.Tag
	db.Where("id IN (?)", []int{1, 2}).Find(&tags)

	fmt.Println("Tags", tags)

	// create pointers to the tags
	var tagPointers []*models.Tag
	for i := range tags {
		tagPointers = append(tagPointers, &tags[i])
	}
	fmt.Println("tagPointers", tagPointers)

	// create a new article and associate it with the existing category and tags
	newArticle := models.Article{
		Title:      "My Article 1",
		Content:    "This is a test article. 1",
		CategoryID: category.ID,
		Tags:       tagPointers,
	}

	// save the new article to the database
	db.Create(&newArticle)

	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged in!",
		"data":    newArticle,
	})
	c.JSON(http.StatusOK, gin.H{
		"messages": "...",
	})
}

var articleBody struct {
	Title      string
	Content    string
	CategoryId uint
	TagId      []int
}

func CreateArticle(c *gin.Context) {
	c.Bind(&articleBody)
	//fmt.Println("Title", articleBody.Title)
	//fmt.Println("Tag ID is :", articleBody.TagId)
	var tagId = articleBody.TagId

	//var newTagId = []int{1, 2}
	//fmt.Println("-->", newTagId)
	var tags []models.Tag
	initializers.DB.Where("id IN (?)", tagId).Find(&tags)
	var tagPointers []*models.Tag
	for i := range tags {
		tagPointers = append(tagPointers, &tags[i])
	}

	newArticle := models.Article{
		Title:      articleBody.Title,
		Content:    articleBody.Content,
		CategoryID: articleBody.CategoryId,
		Tags:       tagPointers,
	}

	// save the new article to the database
	initializers.DB.Create(&newArticle)

	//if result != nil {
	//	fmt.Println(tags)
	//	return
	//}

	//fmt.Println("Tags", tags)

	c.JSON(http.StatusOK, gin.H{
		"data": newArticle,
	})
}

func FetchArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	if err := initializers.DB.Preload("Category").Preload("Tags").First(&article, id).Error; err != nil {
		// handle error
	}

	for _, tag := range article.Tags {
		fmt.Println(tag.ID)
		fmt.Println(tag.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}

func FetchArticleByTag(c *gin.Context) {
	id := c.Param("id")

	var articles []models.Article
	if err := initializers.DB.Joins("JOIN article_tags ON articles.id = article_tags.article_id").
		Where("article_tags.tag_id = ?", id).
		Find(&articles).Error; err != nil {
		// handle error
	}

	//for _, article := range articles {
	//	fmt.Println(article.Title)
	//}

	c.JSON(http.StatusOK, gin.H{
		"data": articles,
	})

}

func FetchArticleByCategory(c *gin.Context) {
	id := c.Param("id")
	var articles []models.Article
	if err := initializers.DB.Joins("JOIN categories ON categories.id = articles.category_id").
		Where("categories.id = ?", id).
		Find(&articles).Error; err != nil {
		// handle error
	}

	//for _, article := range articles {
	//	fmt.Println(article.Title)
	//}

	c.JSON(http.StatusOK, gin.H{
		"data": articles,
	})

}
