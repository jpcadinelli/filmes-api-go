package models

import "gorm.io/gorm"

type Filme struct {
	gorm.Model
	Titulo                  string `json:"titulo"`
	Ano                     string `json:"ano"`
	Sinopse                 string `json:"sinopse"`
	Generos                 string `json:"generos"`
	ClassificacaoIndicativa string `json:"classificacao_indicativa"`
	Classificacao           string `json:"classificacao"`
}

var Filmes []Filme
