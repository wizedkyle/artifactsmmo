package artifacts

import (
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"net/http"
	"os"
	"time"
)

var (
	Client = artifacts{}
)

type artifacts struct {
	Client *http.Client
	Token  *string
}

func Init() {
	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		utils.Logger.Fatal("TOKEN environment variable not set")
	}
	Client.Token = &token
	Client.Client = &http.Client{
		Timeout: time.Second * 10,
	}
}

func (a *artifacts) generateRequest(method string, characterName string, endpoint string, body interface{}) (*http.Request, error) {

	if body == nil {

	} else {

	}
	req, err := http.NewRequest(method, "")
	req.Header.Set("")
}
