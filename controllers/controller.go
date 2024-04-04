package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpcadinelli/filmes-api-go/database"
	"github.com/jpcadinelli/filmes-api-go/models"
)

func ExibeFilmes(c *gin.Context) {
	var filmes []models.Filme
	database.DB.Find(&filmes)
	c.JSON(200, filmes)
}

func CriaNovoFilme(c *gin.Context) {
	var filme models.Filme
	if err := c.ShouldBindJSON(&filme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error()})
		return
	}
	database.DB.Create((&filme))
	c.JSON(http.StatusOK, filme)
}

func BuscaFilmePorId(c *gin.Context) {
	var filme models.Filme
	id := c.Params.ByName("id")
	database.DB.First(&filme, id)
	c.JSON(http.StatusOK, filme)
}
