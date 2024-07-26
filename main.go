package main

import (
	"github.com/rodrigosrf/api-product-go-lab/api"
	"github.com/rodrigosrf/api-product-go-lab/config"
)

// @title           Swagger Product API
// @version         1.0
// @description     Descrição dos recursos da API de produtos
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	cfg := config.LoadConfig()
	db := config.InitDatabase(cfg.ServerAddress)

	api.RegisterProductRoutes(db)
}
