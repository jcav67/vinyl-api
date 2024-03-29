package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumResponse struct {
	Message string  `json:"message"`
	Results []album `json:"result"`
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Naturaleza Sangre", Artist: "Fito Páez", Price: 24.00},
	{ID: "2", Title: "Gaia 2: La voz dormida", Artist: "Mago de Oz", Price: 22.00},
	{ID: "3", Title: "Paraiso AA", Artist: "La Doble A", Price: 14.00},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, AlbumResponse{Message: "", Results: albums})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, AlbumResponse{Message: "Error in body please check", Results: nil})
		//return
	}
	newAlbum.ID = strconv.Itoa(len(albums) + 1)
	albums = append(albums, newAlbum)
	c.JSON(http.StatusOK, AlbumResponse{Message: "", Results: albums})
}

func main() {
	fmt.Println("Inicio de la App")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
