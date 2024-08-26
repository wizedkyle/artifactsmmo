package artifacts

import (
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
)

// FindBuilding
// Returns the coordinates for different buildings.
func (a *artifacts) FindBuilding(building string) (int, int) {
	switch building {
	case models.Bank:
		return models.BankX, models.BankY
	case models.SouthBank:
		return models.SouthBankX, models.SouthBankY
	case models.CookingWorkshop:
		return models.CookingWorkshopX, models.CookingWorkshopY
	case models.GearcraftingWorkshop:
		return models.GearcraftingWorkshopX, models.GearcraftingWorkshopY
	case models.JewelrycraftingWorkshop:
		return models.JewelrycraftingWorkshopX, models.JewelrycraftingWorkshopY
	case models.MiningWorkshop:
		return models.MiningWorkshopX, models.MiningWorkshopY
	case models.WeaponcraftingWorkshop:
		return models.WeaponcraftingWorkshopX, models.WeaponcraftingWorkshopY
	case models.WoodcuttingWorkshop:
		return models.WoodcuttingWorkshopX, models.WoodcuttingWorkshopY
	default:
		return 0, 0
	}
}

// FindRocks
// Returns the coordinates for different rock resources.
func (a *artifacts) FindRocks(rock string) (int, int) {
	if rock == models.Copper {
		return models.CopperX, models.CopperY
	} else if rock == models.Coal {
		return models.CoalX, models.CoalY
	} else if rock == models.Iron {
		return models.IronX, models.IronY
	} else if rock == models.Gold {
		return models.GoldX, models.GoldY
	} else {
		return 0, 0
	}
}

// FindMonster
// Returns the coordinates for different monsters.
func (a *artifacts) FindMonster(monster string) (int, int) {
	switch monster {
	case "chicken":
		return 0, 1
	case "yellow_slime":
		return 4, -1
	case "blue_slime":
		return 2, -1
	case "red_slime":
		return 1, -1
	case "cow":
		return 0, 2
	default:
		return 0, 0
	}
}

// FindTrees
// Returns the coordinates for different tree resources.
func (a *artifacts) FindTrees(tree string) (int, int) {
	return 0, 0
}
