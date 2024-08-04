package artifacts

import (
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"io"
	"net/http"
)

func (a *artifacts) ActionMove(character string, location models.ActionMove) (*models.CharacterMovementResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/move", location)
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		var response models.CharacterMovementResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 403:
		return nil, utils.ErrNotAuthenticated
	case 404:
		return nil, utils.ErrMapNotFound
	case 486:
		return nil, utils.ErrCharacterLockedActionInProgress
	case 490:
		return nil, utils.ErrCharacterAtDestination
	case 498:
		return nil, utils.ErrCharacterNotFound
	case 499:
		return nil, utils.ErrCharacterCooldown
	default:
		fmt.Println(string(body))
		return nil, utils.ErrGenericError
	}
}

func (a *artifacts) ActionGathering(character string) (*models.CharacterGatheringResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/gathering", nil)
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
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
