package controllers

import (
	"net/http"

	"github.com/rg-km/final-project-engineering-40/models"
	"github.com/rg-km/final-project-engineering-40/utils/token"

	"github.com/gin-gonic/gin"
)

func CreateMahasiswa(c *gin.Context) {

	var input models.Mahasiswa

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Mahasiswa{}

	u.Name = input.Name
	u.Nim = input.Nim
	u.Email = input.Email
	u.Semester = input.Semester
	u.Status = input.Status
	u.UniversityID = input.UniversityID
	u.FakultasID = input.FakultasID
	u.ProdiID = input.ProdiID

	_, err := u.SaveMahasiswa()

	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "foreign key constraint failed",
			"error":       err.Error()})

		return
	}

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "registration success"})
}

//get all mahasiswa
func GetAllMahasiswa(c *gin.Context) {

	var mahasiswa []models.Mahasiswa

	err := models.DB.Find(&mahasiswa).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": mahasiswa})
}

// login mahasiswa

func LoginMahasiswa(c *gin.Context) {

	var input models.Mahasiswa

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Mahasiswa{}

	err := models.DB.Where("email = ?", input.Email).Take(&u).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found!"})
		return
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success login",
		"data":        u,
		"token":       token})
}

//delete mahasiswa

func DeleteMahasiswa(c *gin.Context) {

	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	err := models.DB.Where("id = ?", id).Take(&mahasiswa).Delete(&mahasiswa).Error

	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Mahasiswa not found!"})

		return
	}

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "delete mahasiswa success", "data": mahasiswa})
}
