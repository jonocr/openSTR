package controllers

import (
	"fmt"
	"net/http"
	"os"
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
	fmt.Printf("body struct %+v \n", user)
	user := models.User{Email: user.Email, Password: string(hash), Name: user.Name, LastName: user.LastName, IsVerified: false}

	result := database.DB.Create(&user)

	fmt.Printf("user created %+v \n", user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Please check your email and verify your account.",
	})
}

func Singin(c *gin.Context) {

}

func SendVerificationEmail() error {

	emails := []string{os.Getenv("FROM_EMAIL")}
	msg := "3nd email confirmation test"
	subject := "Golang Verification Email from Dev"

	message := "Subject: " + subject + "\n" + msg

	return utils.SendMail(emails, message)

}
