package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/artifactsmmo/v2/internal/build"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "artifacts",
		Short:            "CLI client for Artifacts",
		Long:             "Go based CLI client for interacting with Artifacts MMO (https://artifactsmmo.com/).",
		TraverseChildren: true,
		Version:          build.GetVersion(),
	}
	utils.LoggerInit()
	cmd.AddCommand(NewCmdConfigure())
	cmd.AddCommand(NewCmdMove())
	cmd.AddCommand(NewCmdGather())
	cmd.AddCommand(NewCmdCharacter())
	return cmd
}
