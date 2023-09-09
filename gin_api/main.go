package main

import (
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
	router.Run("localhost:8080")
}

// ptr to context
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
