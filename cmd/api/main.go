package main

import (
	"github.com/wizedkyle/artifactsmmo/v2/internal/database"
	"github.com/wizedkyle/artifactsmmo/v2/internal/routes"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
)

func main() {
	utils.LoggerInit()
	database.Init()
	router := routes.Init()
	routes.GenerateItemRoutes(router)
	routes.GenerateTaskRoutes(router)
	err := router.Run(":9000")
	if err != nil {
		utils.Logger.Fatal("failed to start gin server", zap.Error(err))
	}
}
