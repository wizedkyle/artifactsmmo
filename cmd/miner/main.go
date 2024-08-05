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
		miningResource, ok := os.LookupEnv("MINING_RESOURCE")
		if !ok {
			c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
			if err != nil {
				utils.Logger.Fatal("failed to get character information", zap.Error(err))
			}
			if c.Data.MiningLevel < models.IronLevel {
				x = models.CopperX
				y = models.CopperY
			} else if c.Data.MiningLevel < models.CoalLevel {
				x = models.IronX
				y = models.IronY
			} else if c.Data.MiningLevel < models.GoldLevel {
				x = models.CoalX
				y = models.CoalY
			} else {
				x = models.GoldX
				y = models.GoldY
			}
		} else {
			switch miningResource {
			case models.Copper:
				x = models.CopperX
				y = models.CopperY
			case models.Iron:
				x = models.IronX
				y = models.IronY
			case models.Coal:
				x = models.CoalX
				y = models.CoalY
			case models.Gold:
				x = models.GoldX
				y = models.GoldY
			default:
				x = models.CopperX
				y = models.CopperY
			}
		}
		c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
		if err != nil {
			utils.Logger.Fatal("failed to get character information", zap.Error(err))
		}
		if c.Data.X != x || c.Data.Y != y {
			fmt.Printf("moving character to x=%d y=%d\n", x, y)
			resp, err := artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
				X: x,
				Y: y,
			})
			if err != nil {
				utils.Logger.Error("failed to move character", zap.Error(err))
				return
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
				return
			}
			fmt.Printf("moving character to bank (x=%d y=%d)\n", bankX, bankY)
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
			c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
			if err != nil {
				utils.Logger.Fatal("failed to get character information", zap.Error(err))
			}
			controllers.DepositAllInventory(c.Data.Inventory)
			resp, err = artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
				X: x,
				Y: y,
			})
			if err != nil {
				utils.Logger.Error("failed to move character", zap.Error(err))
				return
			}
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		} else if err == nil {
			fmt.Printf("%s collected %v\n", *artifacts.Client.CharacterName, resp.Data.Details.Items)
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		}
	}
}
