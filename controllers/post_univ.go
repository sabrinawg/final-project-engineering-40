package controllers

import (
	"net/http"

	"github.com/rg-km/final-project-engineering-40/models"

	"github.com/gin-gonic/gin"
)

func GetAllPost(c *gin.Context) {
	var post []models.Post
	err := models.DB.Find(&post).Error
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "failed to get all post",
			"error":       err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success get all post",
		"data":        post})
}

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Post{}

	u.Image = post.Image
	u.Caption = post.Caption
	u.UnivID = post.UnivID

	_, err := u.SavePost()

	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "foreign key constraint failed",
			"error":       err.Error()})

		return
	}

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "create post success"})
}
