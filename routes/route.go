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
	r.Run()
}
