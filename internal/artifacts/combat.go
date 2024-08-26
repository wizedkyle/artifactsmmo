package artifacts

import (
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"io"
	"net/http"
)

// ActionFight
// Instructs the character to fight the monster at the current location.
func (a *artifacts) ActionFight(character string) (*models.CharacterFightData, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/fight", nil)
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
		var response models.CharacterFightData
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 486:
		return nil, utils.ErrCharacterLockedActionInProgress
	case 497:
		return nil, utils.ErrCharacterInventoryFull
	case 498:
		return nil, utils.ErrCharacterNotFound
	case 499:
		return nil, utils.ErrCharacterCooldown
	case 598:
		return nil, utils.ErrMonsterNotFound
	default:
		fmt.Println(string(body))
		return nil, utils.ErrGenericError
	}
}
