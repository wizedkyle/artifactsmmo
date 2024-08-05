package controllers

import (
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
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
