package database

import (
	"log"

	"github.com/jpcadinelli/filmes-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=user password=changeme dbname=filmes-db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banc ode dados")
	}
	DB.AutoMigrate(&models.Filme{})
}
