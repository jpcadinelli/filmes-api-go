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
	if idade == "Livre" {
		database.DB.Find(&filmes, "classificacao_indicativa = 'Livre'")
	} else {
		database.DB.Find(&filmes, "classificacao_indicativa <= ? OR classificacao_indicativa = 'Livre'", idade)
	}
	c.JSON(http.StatusOK, filmes)
}

func BuscaFilmePorGenero(c *gin.Context) {
	var filmes []models.Filme
	genero := c.Params.ByName("genero")
	database.DB.Find(&filmes, "generos LIKE ?", "%"+genero+"%")
	c.JSON(http.StatusOK, filmes)
}

func BuscaFilmePorNota(c *gin.Context) {
	var filmes []models.Filme
	nota := c.Params.ByName("nota")
	database.DB.Order("classificacao desc, titulo").Find(&filmes, "classificacao >= ?", nota)
	c.JSON(http.StatusOK, filmes)
}

func BuscaFilmePorTitulo(c *gin.Context) {
	var filmes []models.Filme
	titulo := c.Params.ByName("titulo")
	database.DB.Order("ano asc").Find(&filmes, "titulo LIKE ?", "%"+titulo+"%")
	c.JSON(http.StatusOK, filmes)
}
