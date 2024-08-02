package main

import (
	"fmt"
	"log"
	"net/http"
	"server/controllers"
	"server/database"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func init() {
	fmt.Println("This will get called on main initialization")
	database.ConnectDatabase()
	database.SyncDatabase()
}

func main() {

	router := gin.Default()
	router.POST("/singup", controllers.Singup)
	router.GET("/verify", controllers.VerifyEmail)
	router.GET("/email", sendEmail)

	router.GET("/albums", getAlbums)
	router.GET("/users", getUsers)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func sendEmail(c *gin.Context) {

	err := controllers.SendVerificationEmail("jono.calvo@gmail.com", "jono", "kja234sdhfkjlahsdflkjasdhflkajhsdf")

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to send verification email.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent",
	})
}

func getUsers(c *gin.Context) {
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
