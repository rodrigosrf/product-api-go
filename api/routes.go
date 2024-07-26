package api

import (
	"github.com/gin-gonic/gin"
	docs "github.com/rodrigosrf/api-product-go-lab/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterProductRoutes(db *mongo.Database) {
	r := gin.Default()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	InitializeHandler(db)

	r.POST("/api/v1/product", CreateProduct)
	r.GET("/api/v1/product", GetProducts)
	r.GET("/products/:id", GetProductByID)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}
