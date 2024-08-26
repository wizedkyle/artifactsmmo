package controllers

import (
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/database"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
	"time"
)

func CompleteCraftingOrder(task models.Task) (string, error) {
	var (
		totalCraftingTrips int
	)
	err := database.Client.UpdateTask(task.Id, "", models.TaskStatusRunning)
	if err != nil {
		utils.Logger.Error("failed to update task status", zap.String("task", task.Id), zap.Error(err))
		return "failed to update task status", err
	}
	item, err := database.Client.GetItem(task.Item)
	if err != nil {
		utils.Logger.Error("failed to get item", zap.Error(err))
		return "failed to get item", err
	}
	c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
	if err != nil {
		utils.Logger.Error("failed to get character information", zap.Error(err))
		return "failed to get character information", err
	}
	err = CraftingLevelCheck(task.Action, item)
	if err != nil {
		return "user crafting level isn't at the right level", err
	}
	err = Move(models.Bank)
	if err != nil {
		return "failed to move character to bank", err
	}
	craftableItems, err := WithdrawCraftingItems(task.Quantity, item)
	if err != nil {
		utils.Logger.Error("failed to withdraw items", zap.Error(err))
	}
	totalCraftingTrips = task.Quantity / craftableItems
	for i := 0; i < totalCraftingTrips; i++ {
		c, err = artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
		if err != nil {
			return "failed to get character information", err
		}
		err = Move(task.Action)
		if err != nil {
			return "failed to move character", err
		}
		craftingResp, err := artifacts.Client.ActionCrafting(*artifacts.Client.CharacterName, models.Item{
			Code:     task.Item,
			Quantity: craftableItems,
		})
		if err != nil {
			utils.Logger.Error("failed to craft item", zap.Error(err))
			return "failed to craft item " + task.Item, err
		}
		time.Sleep(utils.CalculateTimeDifference(craftingResp.Data.Cooldown.StartedAt, craftingResp.Data.Cooldown.Expiration))
		err = Move(models.Bank)
		if err != nil {
			return "failed to move character", err
		}
		c, err = artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
		if err != nil {
			return "failed to get character information", err
		}
		DepositAllInventory(c.Data.Inventory)
		_, err = WithdrawCraftingItems(task.Quantity, item)
		if err != nil {
			return "failed to withdraw items", err
		}
	}
	return "", nil
}
