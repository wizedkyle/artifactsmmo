package artifacts

import (
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"io"
	"net/http"
	"strconv"
)

// GetCharacter
// Returns information about the specified character.
func (a *artifacts) GetCharacter(character string) (*models.CharacterResponse, error) {
	req, err := a.generateRequest(http.MethodGet, "/characters/"+character, nil)
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
		var response models.CharacterResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 403:
		return nil, utils.ErrNotAuthenticated
	case 404:
		return nil, utils.ErrCharacterNotFound
	default:
		fmt.Println(string(body))
		return nil, utils.ErrGenericError
	}
}

// CreateCharacter
// Creates a new character.
func (a *artifacts) CreateCharacter(character models.CreateCharacter) (*models.CharacterResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/characters/create", character)
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
		var response models.CharacterResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 403:
		return nil, utils.ErrNotAuthenticated
	case 494:
		return nil, utils.ErrCharacterNameInUse
	case 495:
		return nil, utils.ErrCharacterLimitReached
	default:
		fmt.Println(string(body))
		return nil, utils.ErrGenericError
	}
}

// ListBankInventory
// Returns all items in the bank.
func (a *artifacts) ListBankInventory(params models.BankInventoryParams) (*models.BankInventoryResponse, error) {
	req, err := a.generateRequest(http.MethodGet, "/my/bank/items", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if params.ItemCode != "" {
		q.Add("item_code", params.ItemCode)
	}
	if params.Page != 0 {
		q.Add("page", strconv.Itoa(params.Page))
	}
	if params.Size != 0 {
		q.Add("size", strconv.Itoa(params.Size))
	}
	req.URL.RawQuery = q.Encode()
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
		var response models.BankInventoryResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 403:
		return nil, utils.ErrNotAuthenticated
	case 404:
		return nil, utils.ErrItemNotFound
	default:
		fmt.Println(string(body))
		return nil, utils.ErrGenericError
	}
}

// ActionMove
// Instructs the character to move to the specified location.
func (a *artifacts) ActionMove(character string, location models.ActionMove) (*models.CharacterMovementResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/move", location)
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
