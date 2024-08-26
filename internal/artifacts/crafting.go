package artifacts

import (
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"io"
	"net/http"
)

// ActionCrafting
// Instructs the character to create the specified item.
func (a *artifacts) ActionCrafting(character string, item models.Item) (*models.CharacterGatheringResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/crafting", item)
	if err != nil {
		return nil, err
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		var response models.CharacterGatheringResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 403:
		return nil, utils.ErrNotAuthenticated
	case 404:
		return nil, utils.ErrItemNotFound
	case 478:
		return nil, utils.ErrItemMissingOrInsufficientQuantity
	case 486:
		return nil, utils.ErrCharacterLockedActionInProgress
	case 493:
		return nil, utils.ErrCharacterNotAtSkillLevel
	case 497:
		return nil, utils.ErrCharacterInventoryFull
	case 498:
		return nil, utils.ErrCharacterNotFound
	case 499:
		return nil, utils.ErrCharacterCooldown
	case 598:
		return nil, utils.ErrResourceNotFound
	default:
		fmt.Println(string(body))
		return nil, utils.ErrGenericError
	}
}

// ActionRecycling
// Instructs the character to recycle the specified item.
func (a *artifacts) ActionRecycling(character string, item models.Item) (*models.CharacterRecyclingResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/recycling", item)
	if err != nil {
		return nil, err
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		var response models.CharacterRecyclingResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 403:
		return nil, utils.ErrNotAuthenticated
	case 404:
		return nil, utils.ErrItemNotFound
	case 473:
		return nil, utils.ErrItemCannotBeRecycled
	case 478:
		return nil, utils.ErrItemMissingOrInsufficientQuantity
	case 486:
		return nil, utils.ErrCharacterLockedActionInProgress
	case 493:
		return nil, utils.ErrCharacterNotAtSkillLevel
	case 497:
		return nil, utils.ErrCharacterInventoryFull
	case 498:
		return nil, utils.ErrCharacterNotFound
	case 499:
		return nil, utils.ErrCharacterCooldown
	case 598:
		return nil, utils.ErrResourceNotFound
	default:
		fmt.Println(string(body))
		return nil, utils.ErrGenericError
	}
}
