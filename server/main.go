package main

import (
	"fmt"
	"net/http"
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

func main() {

	router := gin.Default()
	database.ConnectDatabase()
	router.GET("/albums", getAlbums)
	router.GET("/users", getUsers)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getUsers(c *gin.Context) {

	// connStr := "user=dev dbname=pqgotest sslmode=verify-full"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	age := 21
	rows, err := database.Db.Query("SELECT name FROM users WHERE age = $1", age)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, "Couldn't create the new user.")
	} else {
		c.IndentedJSON(http.StatusOK, rows)
		// ctx.JSON(http.StatusOK, "User is successfully created.")
	}
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
