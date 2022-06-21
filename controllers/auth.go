package controllers

import (
	"net/http"

	"github.com/rg-km/final-project-engineering-40/models"
	"github.com/rg-km/final-project-engineering-40/utils/token"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status code": 200,
		"message":     "Api is working"})
}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	//if isverified false thesn can't login

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	err := models.DB.Where("email = ?", input.Email).Take(&u).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found!"})
		return
	}

	err = models.VerifyPassword(input.Password, u.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Wrong password!"})
		return
	}

	if u.Isverified == false {
		c.JSON(400, gin.H{
			"status code": 400,
			"error":       "Your account is not verified yet!"})
		return
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "Your Token": token})
}

type RegisterInput struct {
	NameUniversity string `json:"name_university"`
	NameRektor     string `json:"name_rektor" binding:"required"`
	KtpRektor      string `json:"ktp_rektor" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.NameUniversity = input.NameUniversity
	u.NameRektor = input.NameRektor
	u.KtpRektor = input.KtpRektor
	u.Email = input.Email
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "registration success"})
}

// update user
func UpdateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Profilepic = input.Profilepic
	u.Bio = input.Bio
	u.Link = input.Link
	u.Whatsapp = input.Whatsapp
	u.UserType = input.UserType
	u.Alamat = input.Alamat

	err = models.DB.Save(&u).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
