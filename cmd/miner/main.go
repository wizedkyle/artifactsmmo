package main

import (
	"errors"
	"fmt"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
)

func main() {
	utils.LoggerInit()
	artifacts.Init()
	resp, err := artifacts.Client.ActionMove("wizedkyle", models.ActionMove{
		X: 3,
		Y: 0,
	})
	if errors.Is(err, utils.ErrCharacterAtDestination) {
		fmt.Println("we are already at the location")
	} else if err != nil {
		utils.Logger.Error("failed to move character", zap.Error(err))
	}
	fmt.Println(resp)
}
