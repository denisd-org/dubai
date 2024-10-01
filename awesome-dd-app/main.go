package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/hello-world", getHelloWorld)
    router.GET("/greet-arnold", getGreetArnold)
    router.GET("/status", getStatus)

    router.Run(":8080")
}

type helloWorld struct {
    Body string `json:"body"`
}

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

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getHelloWorld(c *gin.Context) {
    hw := helloWorld{Body: "awesome-dd-app"}
    c.IndentedJSON(http.StatusOK, hw)
}

func getGreetArnold(c *gin.Context) {
    hw := helloWorld{Body: "Hey Argold!"}
    c.IndentedJSON(http.StatusOK, hw)
}

func getStatus(c *gin.Context) {
    c.Status(http.StatusOK)
}

