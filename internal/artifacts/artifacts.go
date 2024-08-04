package artifacts

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	Client = artifacts{}
)

type artifacts struct {
	CharacterName *string
	Client        *http.Client
	Token         *string
}

const (
	server = "https://api.artifactsmmo.com"
)

func Init() {
	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		viper.SetConfigName("creds")
		viper.AddConfigPath(filepath.Dir(utils.ConfigPath(true)))
		err := viper.ReadInConfig()
		if err != nil {
			utils.Logger.Fatal("failed to read creds file", zap.Error(err))
		}
		characterName := viper.GetString("characterName")
		tokenEncrypted := viper.GetString("token")
		tokenDecrypted := utils.DecryptData(tokenEncrypted)
		Client.CharacterName = &characterName
		Client.Token = &tokenDecrypted
	} else {
		Client.Token = &token
	}
	Client.Client = &http.Client{
		Timeout: time.Second * 10,
	}
}

func (a *artifacts) generateRequest(method string, endpoint string, body interface{}) (*http.Request, error) {
	var (
		bodyByte []byte
	)
	if body == nil {
		bodyByte = nil
	} else {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyByte = b
	}
	req, err := http.NewRequest(method, server+endpoint, bytes.NewReader(bodyByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+*a.Token)
	return req, nil
}
