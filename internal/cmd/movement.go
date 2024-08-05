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

func NewCmdMove() *cobra.Command {
	var (
		x int
		y int
	)
	cmd := &cobra.Command{
		Use:   "move",
		Short: "Move character.",
		Long:  "Moves character to a new position on the map",
		Run: func(cmd *cobra.Command, args []string) {
			utils.LoggerInit()
			artifacts.Init()
			moveCharacter(x, y)
		},
	}
	cmd.Flags().IntVarP(&x, "x", "x", 0, "Specify the X axis coordinate.")
	cmd.Flags().IntVarP(&y, "y", "y", 0, "Specify the Y axis coordinate")
	cmd.MarkFlagRequired("x")
	cmd.MarkFlagRequired("y")
	return cmd
}

func moveCharacter(x int, y int) {
	resp, err := artifacts.Client.ActionMove(*artifacts.Client.CharacterName, models.ActionMove{
		X: x,
		Y: y,
	})
	if errors.Is(err, utils.ErrCharacterAtDestination) {
		fmt.Printf("%s is already at the desired location x=%d y=%d\n", *artifacts.Client.CharacterName, x, y)
		return
	} else if errors.Is(err, utils.ErrCharacterCooldown) {
		fmt.Printf("%s is currently on cooldown\n", *artifacts.Client.CharacterName)
		return
	} else if err != nil {
		utils.Logger.Error("failed to move character", zap.Error(err))
		return
	}
	diff := utils.CalculateTimeDifference(resp.Data.Cooldown.StartedAt, resp.Data.Cooldown.Expiration)
	fmt.Printf("%s is currently at location x=%d y=%d. Character cooldown is %f seconds.\n", *artifacts.Client.CharacterName, x, y, diff.Seconds())
}
