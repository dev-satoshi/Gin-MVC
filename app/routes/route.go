package routes

import (
	"Gin-MVC/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/hello-world", controllers.GetHelloWorld)

	return r
}
