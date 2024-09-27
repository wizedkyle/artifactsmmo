package main

import (
	"errors"
	"github.com/wizedkyle/artifactsmmo/v2/internal/artifacts"
	"github.com/wizedkyle/artifactsmmo/v2/internal/controllers"
	"github.com/wizedkyle/artifactsmmo/v2/internal/database"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.uber.org/zap"
	"time"
)

func main() {
	utils.LoggerInit()
	artifacts.Init()
	database.Init()
	for {
		tasks, err := database.Client.ListTasks("", *artifacts.Client.CharacterName, 1, models.TaskStatusPending)
		if errors.Is(err, utils.ErrTasksNotFound) {
			utils.Logger.Info("no tasks found", zap.Error(err))
			time.Sleep(10 * time.Second)
			continue
		} else if err != nil {
			utils.Logger.Error("failed to get tasks from database", zap.Error(err))
			continue
		}
		for _, task := range *tasks {
			switch task.ActionCategory {
			case models.Combat:
				reason, err := controllers.CompleteCombatOrder(task)
				if err != nil {
					utils.Logger.Error("failed to complete combat order", zap.Error(err))
					err = database.Client.UpdateTask(task.Id, reason, models.TaskStatusSuccess)
					if err != nil {
						utils.Logger.Error("failed to update task status", zap.Error(err))
						continue
					}
					continue
				}
				err = database.Client.UpdateTask(task.Id, "", models.TaskStatusSuccess)
				if err != nil {
					utils.Logger.Error("failed to update task status", zap.Error(err))
					continue
				}
				continue
			case models.Crafting:
				reason, err := controllers.CompleteCraftingOrder(task)
				if err != nil {
					utils.Logger.Error("failed to complete crafting order", zap.Error(err))
					err = database.Client.UpdateTask(task.Id, reason, models.TaskStatusSuccess)
					if err != nil {
						utils.Logger.Error("failed to update task status", zap.Error(err))
						continue
					}
					continue
				}
				err = database.Client.UpdateTask(task.Id, "", models.TaskStatusSuccess)
				if err != nil {
					utils.Logger.Error("failed to update task status", zap.Error(err))
					continue
				}
			default:
				utils.Logger.Info("no valid task action", zap.String("action", task.Action))
				err := database.Client.UpdateTask(task.Id, "no valid task action", models.TaskStatusError)
				if err != nil {
					utils.Logger.Error("failed to update task status", zap.Error(err))
					continue
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}
