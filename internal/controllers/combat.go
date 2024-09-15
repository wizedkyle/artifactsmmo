package controllers

import (
	"errors"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/database"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
	"time"
)

func CompleteCombatOrder(task models.Task) (string, error) {
	err := database.Client.UpdateTask(task.Id, "", models.TaskStatusRunning)
	if err != nil {
		utils.Logger.Error("failed to update task status", zap.String("task", task.Id), zap.Error(err))
		return "failed to update task status", err
	}
	for i := 0; i < task.Quantity; i++ {
		var (
			x int
			y int
		)
		x, y = artifacts.Client.FindMonster(task.Monster)
		c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
		if err != nil {
			utils.Logger.Error("failed to get character information", zap.Error(err))
			return "failed to get character information", err
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
		resp, err := artifacts.Client.ActionFight(*artifacts.Client.CharacterName)
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
			DepositAllInventory(c.Data.Inventory)
			resp, err = artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
				X: x,
				Y: y,
			})
			if err != nil {
				utils.Logger.Error("failed to move character", zap.Error(err))
				continue
			}
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		} else if err != nil {
			utils.Logger.Error("failed to fight", zap.Error(err))
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		} else {
			fmt.Printf("%s fought %s and %s. It dropped %v. Fight number %d\n", *artifacts.Client.CharacterName, task.Monster, resp.Data.Fight.Result, resp.Data.Fight.Drops, i)
			time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
		}
	}
	return "", nil
}
