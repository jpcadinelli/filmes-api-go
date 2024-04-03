package main

import (
	"github.com/jpcadinelli/filmes-api-go/models"
	"github.com/jpcadinelli/filmes-api-go/routes"
)

func main() {
	models.Filmes = []models.Filme{
		{
			Titulo: "Godzilla e Kong: O Novo Império",
			Ano: "2024",
			Sinopse: "O poderoso Kong e o temível Godzilla se unem contra uma colossal ameaça mortal escondida no mundo dos humanos, que ameaça a existência de sua espécie e da nossa. Mergulhando profundamente nos mistérios da Ilha da Caveira e nas origens da Terra Oca, o filme irá explorar a antiga batalha de Titãs que ajudou a forjar esses seres extraordinários e os ligou à humanidade para sempre.",
			Generos: "Ação, Ficção científica e Aventura",
			ClassificacaoIndicativa: "13",
			Classificacao: "67",
		},
		{
			Titulo: "Kung Fu Panda 4",
			Ano: "2024",
			Sinopse: "Po está prestes a se tornar o novo líder espiritual do Vale da Paz, mas antes que possa fazer isso, ele deve encontrar um sucessor para se tornar o novo Dragão Guerreiro. Ele parece encontrar uma em Zhen, uma raposa com muitas habilidades promissoras, mas que não gosta muito da ideia de Po treiná-la.",
			Generos: "Ação, Aventura, Animação, Comédia e Família",
			ClassificacaoIndicativa: "Livre",
			Classificacao: "68",
		},
	}
	routes.HandleRequests()
}
