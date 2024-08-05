package artifacts

import "github.com/wizedkyle/artifactsmmo/v2/internal/models"

// FindBuilding
// Returns the coordinates for different buildings.
func (a *artifacts) FindBuilding(building string) (int, int) {
	if building == models.Bank {
		return models.BankX, models.BankY
	} else {
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
