package main

import (
	"errors"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/controllers"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
	"os"
	"time"
)

func main() {
	utils.LoggerInit()
	artifacts.Init()
	for {
		var (
			x int
			y int
		)
		woodResource, ok := os.LookupEnv("WOOD_RESOURCE")
		if !ok {
			c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
			if err != nil {
				utils.Logger.Error("failed to get character information", zap.Error(err))
				continue
			}
			if c.Data.WoodcuttingLevel < models.SpruceTreeLevel {
				x = models.AshTeeX
				y = models.AshTreeY
			} else if c.Data.WoodcuttingLevel < models.BirchTreeLevel {
				x = models.SpruceTreeX
				y = models.SpruceTreeY
			} else {
				x = models.BirchTreeX
				y = models.BirchTreeY
			}
		} else {
			switch woodResource {
			case models.AshTree:
				x = models.AshTeeX
				y = models.AshTreeY
			case models.SpruceTree:
				x = models.SpruceTreeX
				y = models.SpruceTreeY
			case models.BirchTree:
				x = models.BirchTreeX
				y = models.BirchTreeY
			default:
				x = models.AshTeeX
				y = models.AshTreeY
			}
		}
		c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
		if err != nil {
			utils.Logger.Error("failed to get character information", zap.Error(err))
			continue
		}
		if c.Data.X != x || c.Data.Y != y {
			fmt.Printf("moving character to x=%d y=%d\n", x, y)
			resp, err := artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
				X: x,
				Y: y,
			})
			if err != nil {
				utils.Logger.Error("failed to move character", zap.Error(err))
				continue
			}
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		}
		resp, err := artifacts.Client.ActionGathering(*artifacts.Client.CharacterName)
		if errors.Is(err, utils.ErrCharacterInventoryFull) {
			bankX, bankY := artifacts.Client.FindBuilding(models.Bank)
			resp, err := artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
				X: bankX,
				Y: bankY,
			})
			if err != nil {
				utils.Logger.Error("failed to move character", zap.Error(err))
				continue
			}
			fmt.Printf("moving character to bank (x=%d y=%d)\n", bankX, bankY)
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
			c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
			if err != nil {
				utils.Logger.Error("failed to get character information", zap.Error(err))
				continue
			}
			controllers.DepositAllInventory(c.Data.Inventory)
			resp, err = artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
				X: x,
				Y: y,
			})
			if err != nil {
				utils.Logger.Error("failed to move character", zap.Error(err))
				continue
			}
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		} else if err == nil {
			fmt.Printf("%s collected %v\n", *artifacts.Client.CharacterName, resp.Data.Details.Items)
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		}
	}
}
