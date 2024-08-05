package artifacts

import (
	"encoding/json"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"io"
	"net/http"
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
