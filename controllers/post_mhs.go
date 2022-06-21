package controllers

import (
	"net/http"

	"github.com/rg-km/final-project-engineering-40/models"

	"github.com/gin-gonic/gin"
)

// mahasiswa create post
func CreatePostMahasiswa(c *gin.Context) {

	var input models.Post_mhs

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Post_mhs{}

	u.Image = input.Image
	u.Caption = input.Caption
	u.MhsID = input.MhsID

	_, err := u.SavePostMahasiswa()

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

// get all post
func GetAllPostMahasiswa(c *gin.Context) {

	var post []models.Post_mhs
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
