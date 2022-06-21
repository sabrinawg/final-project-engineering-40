package controllers

import (
	"net/http"

	"github.com/rg-km/final-project-engineering-40/models"

	"github.com/gin-gonic/gin"
)

//Create faculty
func CreateFakultas(c *gin.Context) {
	var fakultas models.Fakultas
	c.BindJSON(&fakultas)
	models.DB.Create(&fakultas)

	err := models.DB.Create(&fakultas).Error
	//error if name faculty nil
	if fakultas.NameFakultas == "" {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "name faculty cannot be empty",
			"error":       err.Error()})
		return
	}
	if fakultas.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "foreign key constraint failed",
			"error":       err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success",
		"data":        fakultas})
}

// get all faculty
func GetAllFakultas(c *gin.Context) {
	var fakultas []models.Fakultas
	models.DB.Find(&fakultas)
	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success",
		"data":        fakultas})
}

// get faculty by query
func GetFakultasByID(c *gin.Context) {
	id := c.Param("id")
	var fakultas models.Fakultas
	models.DB.Where("id = ?", id).Take(&fakultas)

	if fakultas.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Fakultas not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": fakultas})
}

//get faculty by query name
func GetFakultasByName(c *gin.Context) {
	name := c.Param("name_fakultas")
	var fakultas []models.Fakultas
	models.DB.Where("name_fakultas = ?", name).Find(&fakultas)

	if len(fakultas) == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Fakultas not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status code": 200,
		"message":     "success", "data": fakultas})
}

// Update faculty

func UpdateFakultas(c *gin.Context) {
	id := c.Param("id")
	var fakultas models.Fakultas
	models.DB.Where("id = ?", id).Take(&fakultas)

	if fakultas.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Fakultas not found!"})
		return
	}

	if err := c.ShouldBindJSON(&fakultas); err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "no update data",
			"error":       err.Error()})
		return
	}

	c.BindJSON(&fakultas)
	models.DB.Save(&fakultas)

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success update faculty",
		"data":        fakultas})
}

// Delete faculty
func DeleteFakultas(c *gin.Context) {
	id := c.Param("id")
	var fakultas models.Fakultas
	models.DB.Where("id = ?", id).Take(&fakultas)

	if fakultas.ID == 0 {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Fakultas not found!"})
		return
	}

	if err := models.DB.Delete(&fakultas).Error; err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "failed to delete faculty",
			"error":       err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "success delete faculty",
		"data":        fakultas})
}
