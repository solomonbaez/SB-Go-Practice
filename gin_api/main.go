package main

import "fmt"

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
	fmt.Println(albums)
}
