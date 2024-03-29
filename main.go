package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumResponse struct {
	Message string  `json:"message"`
	Results []Album `json:"result"`
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Naturaleza Sangre", Artist: "Fito Páez", Price: 24.00},
	{ID: "2", Title: "Gaia 2: La voz dormida", Artist: "Mago de Oz", Price: 22.00},
	{ID: "3", Title: "Paraiso AA", Artist: "La Doble A", Price: 14.00},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, AlbumResponse{Message: "", Results: albums})
}

func getAlbumsbyId(c *gin.Context) {
	//obtener un query param desde gin
	if id := c.Param("id"); id != "" {
		var foundAlbums []Album
		for _, album := range albums {
			if album.ID == id {

				foundAlbums = append(foundAlbums, album)
				break
			}
		}
		if len(foundAlbums) > 0 {
			c.JSON(http.StatusOK, AlbumResponse{Message: "", Results: foundAlbums})
			return
		} else {
			c.JSON(http.StatusNotFound, AlbumResponse{Message: fmt.Sprintf("Album with id %v not found", id), Results: nil})
			return
		}
	}

	c.JSON(http.StatusBadRequest, AlbumResponse{Message: "Error in request please check", Results: nil})
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

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
	router.GET("/albums/:id", getAlbumsbyId)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
