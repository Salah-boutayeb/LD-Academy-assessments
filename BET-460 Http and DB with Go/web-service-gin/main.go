package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.PUT("/albums/:id", updateAlbumByID)
    router.DELETE("/albums/:id", deleteAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
    // c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

	newAlbum.ID = strconv.Itoa(len(albums)+1)

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
    // c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            // c.IndentedJSON(http.StatusOK, a)
            c.JSON(http.StatusOK, a)
            return
        }
    }
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
    // c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID (c *gin.Context)   {
	id := c.Param("id")
	var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
			a.Artist = newAlbum.Artist
			a.Title = newAlbum.Title
			a.Price = newAlbum.Price
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
	
}

func deleteAlbumByID (c *gin.Context)   {
	id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
			removeAlbum(a.ID)
            c.JSON(http.StatusOK, albums)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
	
}
// this function helps us to remove album from album slice
func removeAlbum(id string)   {
	index,_ := strconv.Atoi(id)
	index-=1
	albums = append(albums[:index],albums[index+1:]...)
	
	
}

