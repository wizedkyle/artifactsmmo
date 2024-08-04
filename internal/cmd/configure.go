package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func NewCmdConfigure() *cobra.Command {
	var (
		characterName string
		token         string
	)
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configures artifactsmmo credentials.",
		Long:  "Configures artifactsmmo cli with a configuration file stating the token and character name for use in subsequent commands.",
		Run: func(cmd *cobra.Command, args []string) {
			setCredentials(characterName, token)
		},
	}
	cmd.Flags().StringVarP(&characterName, "character-name", "c", "", "Specify the character name you want to control.")
	cmd.Flags().StringVarP(&token, "token", "t", "", "Specify the API token for your Artifacts account.")
	cmd.MarkFlagRequired("character-name")
	cmd.MarkFlagRequired("token")
	return cmd
}

func setCredentials(characterName string, token string) {
	var credentials models.Credentials
	credentials.CharacterName = characterName
	credentials.Token = utils.EncryptData(token)
	configFilePath := filepath.Dir(utils.ConfigPath(true))
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err = os.MkdirAll(configFilePath, 0755)
		if err != nil {
			utils.Logger.Fatal("failed to create folder structure for credentials file", zap.Error(err))
		}
	}
	credentialsFile, err := json.MarshalIndent(credentials, "", "    ")
	if err != nil {

	}
	err = os.WriteFile(utils.ConfigPath(true), credentialsFile, 0644)
	if err != nil {
		utils.Logger.Fatal("failed to write credentials to file " + utils.ConfigPath(true))
	} else {
		fmt.Printf("credentials file saved to: %s\n", utils.ConfigPath(true))
		os.Exit(0)
	}
}
