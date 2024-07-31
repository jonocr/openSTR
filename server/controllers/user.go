package controllers

import (
	"net/http"
	"server/database"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func Singup(c *gin.Context) {

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	token, err := utils.GenerateRandomHex(128)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate verification code.",
		})
		return
	}

	user := models.User{Email: user.Email, Password: string(hash), Name: user.Name, LastName: user.LastName, IsVerified: false}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	tokenObj := models.UserToken{
		Token:  token,
		User:   user,
		UserId: user.Id,
	}

	verificationResult := database.DB.Create(&tokenObj)

	if verificationResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create new verification code.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Please check your email and verify your account.",
		"user":    user,
		"token":   token,
	})
}

func Singin(c *gin.Context) {

}

func SendVerificationEmail(email string, verificationCode string) error {
	emails := []string{email}
	//TODO: add a msg and subject to msg config file
	msg := "3nd email confirmation test"
	subject := "Golang Verification Email from Dev"
	message := "Subject: " + subject + "\n" + msg

	return utils.SendMail(emails, subject, message)

}
