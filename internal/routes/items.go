package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wizedkyle/artifactsmmo/v2/internal/database"
	"github.com/wizedkyle/artifactsmmo/v2/internal/middleware"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"net/http"
	"strconv"
)

func GenerateItemRoutes(router *gin.Engine) {
	itemRoutes := router.Group("/v1/items")
	itemRoutes.Use(middleware.ValidateTransactionId())
	{
		itemRoutes.GET("", listItems)
		itemRoutes.GET(":code", getItem)
	}
}

func listItems(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	limit := utils.QueryLimit(c)
	itemType := c.Query("type")
	subType := c.Query("subtype")
	level, err := strconv.Atoi(c.Query("level"))
	if err != nil {
		utilErr := utils.GenerateError(models.ItemRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	results, err := database.Client.ListItems(limit, models.ListItemParameters{
		Type:    itemType,
		SubType: subType,
		Level:   level,
	})
	if errors.Is(err, utils.ErrItemsNotFound) {
		utilErr := utils.GenerateError(models.ItemRetrieved, utils.ItemsNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(models.ItemRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func getItem(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	code := c.Param("code")
	result, err := database.Client.GetItem(code)
	if errors.Is(err, utils.ErrItemNotFound) {
		utilErr := utils.GenerateError(models.ItemRetrieved, utils.ItemNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(models.ItemRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}
