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
	// TODO: add the rest of the workshops
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
	default:
		return errors.New("unknown action " + craftingAction)
	}
}
