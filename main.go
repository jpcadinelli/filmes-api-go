package main

import (
	"github.com/jpcadinelli/filmes-api-go/database"
	"github.com/jpcadinelli/filmes-api-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
