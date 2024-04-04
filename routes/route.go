package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpcadinelli/filmes-api-go/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/filmes", controllers.ExibeFilmes)
	r.POST("/filmes", controllers.CriaNovoFilme)
	r.GET("/filmes/:id", controllers.BuscaFilmePorId)
	r.DELETE("/filmes/:id", controllers.DeletaFilme)
	r.PATCH("/filmes/:id", controllers.EditaFilme)
	r.GET("/filmes/ci/:idade", controllers.BuscaFilmePorIdade)
	r.GET("/filmes/genero/:genero", controllers.BuscaFilmePorGenero)
	r.GET("/filmes/nota/:nota", controllers.BuscaFilmePorNota)
	r.Run()
}
