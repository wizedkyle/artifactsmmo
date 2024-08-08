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
		fishingResource, ok := os.LookupEnv("FISHING_RESOURCE")
		if !ok {
			c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
			if err != nil {
				utils.Logger.Error("failed to get character information", zap.Error(err))
				continue
			}
			if c.Data.FishingLevel < models.ShrimpLevel {
				x = models.GudgeonX
				y = models.GudgeonY
			} else if c.Data.MiningLevel < models.TroutLevel {
				x = models.ShrimpX
				y = models.ShrimpY
			} else if c.Data.MiningLevel < models.BassLevel {
				x = models.TroutX
				y = models.TroutY
			} else {
				x = models.BassX
				y = models.BassY
			}
		} else {
			switch fishingResource {
			case models.Gudgeon:
				x = models.GudgeonX
				y = models.GudgeonY
			case models.Shrimp:
				x = models.ShrimpX
				y = models.ShrimpY
			case models.Trout:
				x = models.TroutX
				y = models.TroutY
			case models.Bass:
				x = models.BassX
				y = models.BassY
			default:
				x = models.GudgeonX
				y = models.GudgeonY
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
