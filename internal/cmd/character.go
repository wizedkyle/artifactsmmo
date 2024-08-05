package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
)

func NewCmdCharacter() *cobra.Command {
	var (
		name   string
		men1   bool
		men2   bool
		men3   bool
		women1 bool
		women2 bool
		women3 bool
		skin   string
	)
	cmd := &cobra.Command{
		Use:   "create-character",
		Short: "Creates a new character for the account",
		Long:  "Creates a new character for the account",
		Run: func(cmd *cobra.Command, args []string) {
			utils.LoggerInit()
			artifacts.Init()
			if men1 {
				skin = "men1"
			} else if men2 {
				skin = "men2"
			} else if men3 {
				skin = "men3"
			} else if women1 {
				skin = "women1"
			} else if women2 {
				skin = "women2"
			} else if women3 {
				skin = "women3"
			}
			createCharacter(name, skin)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the character")
	cmd.Flags().BoolVar(&men1, "men1", false, "Create a character with men1 skin")
	cmd.Flags().BoolVar(&men2, "men2", false, "Create a character with men2 skin")
	cmd.Flags().BoolVar(&men3, "men3", false, "Create a character with men3 skin")
	cmd.Flags().BoolVar(&women1, "women1", false, "Create a character with women1 skin")
	cmd.Flags().BoolVar(&women2, "women2", false, "Create a character with women2 skin")
	cmd.Flags().BoolVar(&women3, "women3", false, "Create a character with women3 skin")
	cmd.MarkFlagRequired("name")
	return cmd
}

func createCharacter(name string, skin string) {
	res, err := artifacts.Client.CreateCharacter(models.CreateCharacter{
		Name: name,
		Skin: skin,
	})
	if errors.Is(err, utils.ErrNotAuthenticated) {
		fmt.Println(utils.NotAuthenticated)
		return
	} else if errors.Is(err, utils.ErrCharacterNameInUse) {
		fmt.Printf(utils.CharacterNameInUse)
		return
	} else if errors.Is(err, utils.ErrCharacterLimitReached) {
		fmt.Println(utils.CharacterLimitReached)
		return
	} else if err != nil {
		utils.Logger.Error("failed to create character", zap.Error(err))
	}
	fmt.Printf("%s created.", res.Data.Name)
}
