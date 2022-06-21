package controllers

import (
	"net/http"

	"github.com/rg-km/final-project-engineering-40/models"

	"github.com/gin-gonic/gin"
)

//create new prodi
func CreateProdi(c *gin.Context) {
	var prodi models.Prodi
	c.BindJSON(&prodi)
	models.DB.Create(&prodi)

	err := models.DB.Create(&prodi).Error
	//error if name faculty nil
	if prodi.NameProdi == "" {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "name prodi cannot be empty",
			"error":       err.Error()})
		return
	}
	if prodi.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "foreign key constraint failed",
			"error":       err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success",
		"data":        prodi})
}

// get all faculty
func GetAllProdi(c *gin.Context) {
	var prodi []models.Prodi
	models.DB.Find(&prodi)
	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success",
		"data":        prodi})
}

func GetProdiByID(c *gin.Context) {
	id := c.Param("id")
	var prodi models.Prodi
	models.DB.Where("id = ?", id).Take(&prodi)

	if prodi.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Prodi not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status code": 200,
		"message":     "success", "data": prodi})
}

//get faculty by query name
func GetProdiByName(c *gin.Context) {
	name := c.Param("name_prodi")
	var prodi []models.Prodi
	models.DB.Where("name_prodi = ?", name).Find(&prodi)

	if len(prodi) == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Prodi not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status code": 200,
		"message":     "success", "data": prodi})
}

// update prodi

func UpdateProdi(c *gin.Context) {
	id := c.Param("id")
	var prodi models.Prodi
	models.DB.Where("id = ?", id).Take(&prodi)

	if prodi.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Prodi not found!"})
		return
	}

	c.BindJSON(&prodi)
	models.DB.Save(&prodi)

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success update prodi",
		"data":        prodi})
}

// delete prodi

func DeleteProdi(c *gin.Context) {
	id := c.Param("id")
	var prodi models.Prodi
	models.DB.Where("id = ?", id).Take(&prodi)

	if prodi.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Prodi not found!"})
		return
	}

	models.DB.Delete(&prodi)

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success delete prodi",
		"data":        prodi})
}
