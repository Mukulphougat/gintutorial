package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	server := gin.Default()

	server.GET("/hello", func(context *gin.Context) {
		//fmt.Println(context)
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	server.GET("/albums", getAlbums)
	log.Fatal(http.ListenAndServe("localhost:8080", server))
	server.Run()
}
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
