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

func (a *artifacts) ListEvents(params models.GetAllEventsQueryParameters) (*models.EventResponse, error) {
	req, err := a.generateRequest(http.MethodGet, "/events", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if params.Page != 0 {
		q.Add("page", strconv.Itoa(params.Page))
	}
	if params.Size != 0 {
		q.Add("size", strconv.Itoa(params.Size))
	} else {
		q.Add("size", strconv.Itoa(50))
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
		var response models.EventResponse
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
