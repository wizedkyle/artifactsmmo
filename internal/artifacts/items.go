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

// GetItems
// Returns information about the items based on query parameters.
func (a *artifacts) GetItems(params models.GetAllItemsQueryParameters) (*models.ItemsResponse, error) {
	req, err := a.generateRequest(http.MethodGet, "/items/", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if params.CraftMaterial != "" {
		q.Add("craft_material", params.CraftMaterial)
	}
	if params.CraftSkill != "" {
		q.Add("craft_skill", params.CraftSkill)
	}
	if params.MaxLevel != 0 {
		q.Add("max_level", strconv.Itoa(params.MaxLevel))
	}
	if params.MinLevel != 0 {
		q.Add("min_level", strconv.Itoa(params.MinLevel))
	}
	if params.Name != "" {
		q.Add("name", params.Name)
	}
	if params.Page != 0 {
		q.Add("page", strconv.Itoa(params.Page))
	}
	if params.Size != 0 {
		q.Add("size", strconv.Itoa(params.Size))
	} else {
		q.Add("size", strconv.Itoa(50))
	}
	if params.Type != "" {
		q.Add("type", params.Type)
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
		var response models.ItemsResponse
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

// GetItem
// Returns information about a specific item.
func (a *artifacts) GetItem(code string) (*models.ItemResponse, error) {
	req, err := a.generateRequest(http.MethodGet, "/items/"+code, nil)
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
		var response models.ItemResponse
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
