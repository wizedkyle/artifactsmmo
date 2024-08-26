package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wizedkyle/artifactsmmo/v2/internal/database"
	"github.com/wizedkyle/artifactsmmo/v2/internal/middleware"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"net/http"
)

func GenerateTaskRoutes(router *gin.Engine) {
	taskRoutes := router.Group("/v1/tasks")
	taskRoutes.Use(middleware.ValidateTransactionId())
	{
		taskRoutes.POST("", createTask)
		taskRoutes.GET("", listTasks)
		taskRoutes.GET(":id", getTask)
	}
}

func createTask(c *gin.Context) {
	var request models.CreateTask
	transactionId, _ := utils.GetTransactionIdHeader(c)
	if err := c.ShouldBindJSON(&request); err != nil {
		utilErr := utils.GenerateError(models.TaskCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateTask(models.Task{
		Action:         request.Action,
		ActionCategory: request.ActionCategory,
		Monster:        request.Monster,
		Item:           request.Item,
		Quantity:       request.Quantity,
		Character:      request.Character,
	})
	if err != nil {
		utilErr := utils.GenerateError(models.TaskCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func listTasks(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	action := c.Query("action")
	limit := utils.QueryLimit(c)
	status := c.Query("status")
	results, err := database.Client.ListTasks(action, "", limit, status)
	if errors.Is(err, utils.ErrTasksNotFound) {
		utilErr := utils.GenerateError(models.TaskRetrieved, utils.TasksNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(models.TaskRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func getTask(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	id := c.Param("id")
	result, err := database.Client.GetTask(id)
	if errors.Is(err, utils.ErrTaskNotFound) {
		utilErr := utils.GenerateError(models.TaskRetrieved, utils.TaskNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(models.TaskRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}
