package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
)

func NewCmdGather() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gather",
		Short: "Gather resource.",
		Long:  "Gathers resources from the characters current location",
		Run: func(cmd *cobra.Command, args []string) {
			utils.LoggerInit()
			artifacts.Init()
			gatherResource()
		},
	}
	return cmd
}

func gatherResource() {
	res, err := artifacts.Client.ActionGathering(*artifacts.Client.CharacterName)
	if errors.Is(err, utils.ErrCharacterInventoryFull) {
		fmt.Printf("%s inventory is full\n", *artifacts.Client.CharacterName)
		return
	} else if errors.Is(err, utils.ErrCharacterNotAtSkillLevel) {
		fmt.Printf("%s not at the required skill level\n", *artifacts.Client.CharacterName)
		return
	} else if errors.Is(err, utils.ErrResourceNotFound) {
		fmt.Println("no resources found at current location")
		return
	} else if errors.Is(err, utils.ErrCharacterCooldown) {
		fmt.Printf("%s is currently on cooldown\n", *artifacts.Client.CharacterName)
		return
	}
	fmt.Printf("%s gathered %v. Character cooldown is %d\n", *artifacts.Client.CharacterName, res.Data.Details.Items, res.Data.Cooldown.RemainingSeconds)
}
