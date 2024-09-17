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
	switch rock {
	case models.Copper:
		return models.CopperX, models.CopperY
	case models.Iron:
		return models.IronX, models.IronY
	case models.Coal:
		return models.CoalX, models.CoalY
	case models.Gold:
		return models.GoldX, models.GoldY
	default:
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
	case "green_slime":
		return 0, -1
	case "cow":
		return 0, 2
	case "flying_serpent":
		return 5, 4
	case "wolf":
		return -2, 1
	case "skeleton":
		return 8, 6
	default:
		return 0, 0
	}
}

// FindTrees
// Returns the coordinates for different tree resources.
func (a *artifacts) FindTrees(tree string) (int, int) {
	switch tree {
	case models.AshTree:
		return models.AshTeeX, models.AshTreeY
	case models.SpruceTree:
		return models.SpruceTreeX, models.SpruceTreeY
	case models.BirchTree:
		return models.BirchTreeX, models.BirchTreeY
	case models.DeadTree:
		return models.DeadTreeX, models.DeadTreeY
	default:
		return 0, 0
	}
}
