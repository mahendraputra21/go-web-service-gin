package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//album slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	//init event
	router := gin.Default()
	router.GET("/", homeAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PATCH("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
}

func homeAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome home!")
}

// getAlbum responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// deleteAlbumByID locates the album whose ID Value matches the id
// parameter sent by the client, then returns that info
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, singleAlbum := range albums {
		if singleAlbum.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, "The Album with id "+id+" has been deleted successfully")
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// updateAlbumbyID locates the album whose ID value matches the id
// parameter sent by the client, then return that info
func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbums album

	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	json.Unmarshal(reqBody, &updatedAlbums)

	for i, singleAlbum := range albums {
		if singleAlbum.ID == id {
			singleAlbum.Artist = updatedAlbums.Artist
			singleAlbum.Title = updatedAlbums.Title
			singleAlbum.Price = updatedAlbums.Price
			albums = append(albums[:i], singleAlbum)
			c.IndentedJSON(http.StatusOK, "The Album with id "+id+" has been updated successfully")
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
