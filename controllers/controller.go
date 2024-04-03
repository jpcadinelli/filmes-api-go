package controllers

import "github.com/gin-gonic/gin"

func ExibeFilmes(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":                       "1",
		"título":                   "Godzilla e Kong: O Novo Império",
		"ano":                      "2024",
		"sinopse":                  "O poderoso Kong e o temível Godzilla se unem contra uma colossal ameaça mortal escondida no mundo dos humanos, que ameaça a existência de sua espécie e da nossa. Mergulhando profundamente nos mistérios da Ilha da Caveira e nas origens da Terra Oca, o filme irá explorar a antiga batalha de Titãs que ajudou a forjar esses seres extraordinários e os ligou à humanidade para sempre.",
		"generos":                  "Ação, Ficção científica e Aventura",
		"classificação indicativa": "13",
		"classificação":            "67",
	})
}
