package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jpcadinelli/filmes-api-go/models"
)

func ExibeFilmes(c *gin.Context) {
	c.JSON(200, models.Filmes)
}
