package controllers

import (
	"errors"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"strconv"
	"time"
)

func Move(destinationBuildingName string) error {
	x, y := artifacts.Client.FindBuilding(destinationBuildingName)
	c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
	if err != nil {
		return errors.New("failed to get character information")
	}
	if c.Data.X != x || c.Data.Y != y {
		fmt.Printf("moving character to x=%d y=%d\n", x, y)
		resp, err := artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
			X: x,
			Y: y,
		})
		if err != nil {
			return errors.New("failed to move character")
		}
		time.Sleep(utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration))
	}
	return nil
}

// CraftingLevelCheck
// Looks at the item to be crafted and checks that the users crafting level allows it.
func CraftingLevelCheck(craftingAction string, item *models.ItemDetails) error {
	c, err := artifacts.Client.GetCharacter(*artifacts.Client.CharacterName)
	if err != nil {
		return errors.New("failed to get character information")
	}
	switch craftingAction {
	case models.CookingWorkshop:
		if item.Craft.Level > c.Data.CookingLevel {
			utils.Logger.Error("character cooking not at required level")
			return errors.New("cooking not at required level, must be equal to or above " + strconv.Itoa(item.Craft.Level))
		}
		return nil
	case models.GearcraftingWorkshop:
		if item.Craft.Level > c.Data.GearCraftingLevel {
			utils.Logger.Error("character gear crafting not at required level")
			return errors.New("gear crafting not at required level, must be equal to or above " + strconv.Itoa(item.Craft.Level))
		}
		return nil
	case models.JewelrycraftingWorkshop:
		if item.Craft.Level > c.Data.JewelryCraftingLevel {
			utils.Logger.Error("character jewelry crafting not at required level")
			return errors.New("jewel crafting not at required level, must be equal to or above " + strconv.Itoa(item.Craft.Level))
		}
		return nil
	case models.MiningWorkshop:
		if item.Craft.Level > c.Data.MiningLevel {
			utils.Logger.Error("character mining not at required level")
			return errors.New("mining not at required level, must be equal to or above " + strconv.Itoa(item.Craft.Level))
		}
		return nil
	case models.WeaponcraftingWorkshop:
		if item.Craft.Level > c.Data.WeaponCraftingLevel {
			utils.Logger.Error("character weapon crafting not at required level")
			return errors.New("weapon crafting not at required level, must be equal to or above " + strconv.Itoa(item.Craft.Level))
		}
		return nil
	case models.WoodcuttingWorkshop:
		if item.Craft.Level > c.Data.WoodcuttingLevel {
			utils.Logger.Error("character wood crafting not at required level")
			return errors.New("woodcutting not at required level, must be equal to or above " + strconv.Itoa(item.Craft.Level))
		}
		return nil
	default:
		return errors.New("unknown action " + craftingAction)
	}
}
