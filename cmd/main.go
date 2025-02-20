package main

import (
	"log"

	_ "github.com/SolidShake/gin-swagger-example/docs"
	"github.com/SolidShake/gin-swagger-example/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Example API
// @version         1.0
// @description     This is a sample server.
// @host           localhost:8080
// @BasePath       /api/v1
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", h.GetUsers)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
