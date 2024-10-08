package artifacts

import (
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"io"
	"net/http"
)

// ActionWithdrawBank
// Instructs the character to withdraw the specified resources from the bank.
func (a *artifacts) ActionWithdrawBank(character string, withdrawal models.ActionWithdrawBank) (*models.WithdrawBankResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/bank/withdraw", withdrawal)
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
		var response models.WithdrawBankResponse
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

// ActionDepositBank
// Instructs the character to deposit the specified resources in the bank.
func (a *artifacts) ActionDepositBank(character string, deposit models.ActionDepositBank) (*models.ActionDepositBankResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/bank/deposit", deposit)
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
		var response models.ActionDepositBankResponse
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

// ActionGathering
// Instructs the character to gather resources at the current location.
func (a *artifacts) ActionGathering(character string) (*models.CharacterGatheringResponse, error) {
	req, err := a.generateRequest(http.MethodPost, "/my/"+character+"/action/gathering", nil)
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
