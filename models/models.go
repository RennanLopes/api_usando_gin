package models

type Pessoa struct {
	Nome     string  `json:"nome"`
	DataNasc string  `json:"data_nasc"`
	Peso     float64 `json:"peso"`
	Altura   float64 `json:"altura"`
}
