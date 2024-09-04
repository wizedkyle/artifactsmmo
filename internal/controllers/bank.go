package controllers

import (
	"errors"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func DepositAllInventory(inventory []models.InventorySlot) {
	for _, item := range inventory {
		if item.Code == "" && item.Quantity == 0 {
			continue
		} else {
			b, err := artifacts.Client.ActionDepositBank(*artifacts.Client.CharacterName, models.ActionDepositBank{
				Code:     item.Code,
				Quantity: item.Quantity,
			})
			if err != nil {
				utils.Logger.Error("failed to deposit bank", zap.Error(err))
			}
			fmt.Printf("%s deposited %d of %s\n", *artifacts.Client.CharacterName, item.Quantity, item.Code)
			time.Sleep(utils.CalculateTimeDifference(b.Data.Cooldown.StartedAt, b.Data.Cooldown.Expiration))
		}
	}
}

// WithdrawCraftingItems
// Calculates the maximum items that can be crafted per inventory. Retrieves all the materials and returns the maximum craftable per inventory.
func WithdrawCraftingItems(quantity int, items *models.ItemDetails) (int, error) {
	var (
		currentInventorySize   int
		maxCraftableItems      int
		totalResourcesRequired int
	)
	for _, item := range items.Craft.Items {
		totalResourcesRequired = totalResourcesRequired + item.Quantity
	}
	c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
	if err != nil {
		return 0, errors.New("failed to get character information")
	}
	maxCraftableItems = c.Data.InventoryMaxItems / totalResourcesRequired
	if quantity < maxCraftableItems {
		maxCraftableItems = quantity
	}
	if len(c.Data.Inventory) == 0 {
		currentInventorySize = c.Data.InventoryMaxItems
	} else {
		for _, inventory := range c.Data.Inventory {
			currentInventorySize += inventory.Quantity
		}
		spareInventorySize := c.Data.InventoryMaxItems - currentInventorySize
		if spareInventorySize < totalResourcesRequired*maxCraftableItems {
			err := Move(models.Bank)
			if err != nil {
				return 0, err
			}
			DepositAllInventory(c.Data.Inventory)
		}
	}
	for _, item := range items.Craft.Items {
		err := Move(models.Bank)
		if err != nil {
			return 0, err
		}
		resp, err := artifacts.Client.ActionWithdrawBank(*artifacts.Client.CharacterName, models.ActionWithdrawBank{
			Code:     item.Code,
			Quantity: item.Quantity * maxCraftableItems,
		})
		if err != nil {
			utils.Logger.Error("failed to withdraw bank", zap.Error(err))
			return 0, errors.New("failed to withdraw " + strconv.Itoa(item.Quantity*totalResourcesRequired) + " of " + item.Code + " from bank")
		}
		time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
	}
	return maxCraftableItems, nil
}
