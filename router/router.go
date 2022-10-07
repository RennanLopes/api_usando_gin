package router

import (
	"api_usando_gin/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()
	router.GET("/pessoas", controller.GetPessoa)

	router.Run("localhost:8080")

}
