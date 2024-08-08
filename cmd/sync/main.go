package main

import (
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/database"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
)

func main() {
	utils.LoggerInit()
	artifacts.Init()
	database.Init()
	ItemTypes := []string{
		"consumable",
		"body_armor",
		"weapon",
		"resource",
		"leg_armor",
		"helmet",
		"boots",
		"shield",
		"amulet",
		"ring",
	}
	for _, itemType := range ItemTypes {
		err := database.Client.DeleteItems(itemType)
		if err != nil {
			utils.Logger.Error("Failed to delete item type", zap.String("type", itemType), zap.Error(err))
		}
		var items []interface{}
		resp, err := artifacts.Client.GetItems(models.GetAllItemsQueryParameters{
			Type: itemType,
			Size: 100,
		})
		if err != nil {
			utils.Logger.Error("failed to get items", zap.Error(err))
		}
		for _, item := range resp.Data {
			items = append(items, item)
		}
		err = database.Client.CreateItems(items)
		if err != nil {
			utils.Logger.Error("failed to create items", zap.Error(err))
		}
	}
}
