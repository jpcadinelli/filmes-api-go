package models

import "gorm.io/gorm"

type Filme struct {
	gorm.Model
	Titulo                  string `json:"título"`
	Ano                     string `json:"ano"`
	Sinopse                 string `json:"sinopse"`
	Generos                 string `json:"generos"`
	ClassificacaoIndicativa string `json:"classificação indicativa"`
	Classificacao           string `json:"classificação"`
}

var Filmes []Filme
