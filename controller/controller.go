package controller

import (
	"api_usando_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var pessoas = []models.Pessoa{
	{Nome: "João", DataNasc: "01/01/2001", Peso: 62.55, Altura: 1.87},
	{Nome: "Carlos", DataNasc: "04/06/2005", Peso: 80.25, Altura: 1.58},
	{Nome: "Paulo", DataNasc: "22/05/2010", Peso: 56.10, Altura: 1.63},
}

func GetPessoa(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, pessoas)

}
