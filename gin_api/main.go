package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "The Beach", Artist: "Amiture", Price: 9.00},
	{ID: "2", Title: "Careful", Artist: "Boy Harsher", Price: 8.00},
	{ID: "3", Title: "Visions", Artist: "Grimes", Price: 14.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums) // pass function rather than function result
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/delete/:id", deleteAlbumByID)
	router.Run("localhost:8080")
}

// ptr to context
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	// pass ref
	if e := c.BindJSON(&newAlbum); e != nil {
		fmt.Errorf("Failed to POST: %e", e)
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}

	if index >= 0 {
		albums = append(albums[:index], albums[index+1:]...)
		c.Status(http.StatusNoContent)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}
