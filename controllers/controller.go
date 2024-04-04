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
	if filme.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Filme não encontrado"})
		return
	}
	c.JSON(http.StatusOK, filme)
}

func DeletaFilme(c *gin.Context) {
	var filme models.Filme
	id := c.Params.ByName("id")
	database.DB.Delete(&filme, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Filme deletado com sucesso",
	})
}

func EditaFilme(c *gin.Context) {
	var filme models.Filme
	id := c.Params.ByName("id")
	database.DB.First(&filme, id)
	if err := c.ShouldBindJSON(&filme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&filme).UpdateColumns(filme)
	c.JSON(http.StatusOK, filme)
}

func BuscaFilmePorIdade(c *gin.Context) {
	var filmes []models.Filme
	idade := c.Params.ByName("idade")
	database.DB.Find(&filmes, map[string]interface{}{"classificacao_indicativa": idade})
	c.JSON(http.StatusOK, filmes)
}
