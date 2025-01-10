package aplication

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	executionhistoryServices "github.com/gody-server/app/executionhistory/domain/services"
	executionhistoryDatabase "github.com/gody-server/app/executionhistory/infraestructure"

	// Replace for dto
	executionhistoryModel "github.com/gody-server/app/executionhistory/domain/models"
)

// Create a executionhistoryDb instance
var executionhistoryDb = executionhistoryDatabase.NewExecutionhistoryDb()

// Create a Service instance using the executionhistoryDb
var service = executionhistoryServices.NewService(executionhistoryDb)

// Post Executionhistory
// @Summary Post Executionhistory
// @Schemes
// @Description Post Executionhistory
// @Tags Executionhistory
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body executionhistoryModel.Executionhistory true "Body to create Executionhistory"
// @Success 200
// @Router /execution-history [Post]
func CreateExecutionhistory(c *gin.Context) {
	var executionhistory executionhistoryModel.Executionhistory
	if err := c.ShouldBindJSON(&executionhistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateExecutionhistory(executionhistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "executionhistory created",
	})
}

// Get Executionhistory
// @Summary Get Executionhistory
// @Schemes
// @Description Get Executionhistory
// @Tags Executionhistory
// @Security BearerAuth
// @Param actionId path int true "actionId"
// @Accept json
// @Produce json
// @Success 200
// @Router /execution-history/all/by-action/{actionId} [Get]
func GetExecutionhistory(c *gin.Context) {
	actionId, err := strconv.Atoi(c.Param("actionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := service.GetExecutionhistory(actionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Get Executionhistory
// @Summary Get Executionhistory
// @Schemes
// @Description Get Executionhistory
// @Tags Executionhistory
// @Security BearerAuth
// @Param actionId path int true "actionId"
// @Accept json
// @Produce json
// @Success 200
// @Router /execution-history/{actionId} [Get]
func GetOneExecutionhistory(c *gin.Context) {
	actionId, err := strconv.Atoi(c.Param("actionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.GetOneExecutionhistory(actionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Put Executionhistory
// @Summary Put Executionhistory
// @Schemes
// @Description Put Executionhistory
// @Tags Executionhistory
// @Security BearerAuth
// @Param executionHistoryId path int true "executionHistoryId"
// @Accept json
// @Produce json
// @Param Body body executionhistoryModel.Executionhistory true "Body to update Executionhistory"
// @Success 200
// @Router /execution-history/{executionHistoryId} [Put]
func UpdateExecutionhistory(c *gin.Context) {
	var executionhistory executionhistoryModel.Executionhistory
	if err := c.ShouldBindJSON(&executionhistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	executionHistoryId, err := strconv.Atoi(c.Param("executionHistoryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateExecutionhistory(executionHistoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "executionhistory updated",
	})
}

// Delete Executionhistory
// @Summary Delete Executionhistory
// @Schemes
// @Description Delete Executionhistory
// @Tags Executionhistory
// @Security BearerAuth
// @Param executionHistoryId path int true "executionHistoryId"
// @Accept json
// @Produce json
// @Success 200
// @Router /execution-history/{executionHistoryId} [Delete]
func DeleteExecutionhistory(c *gin.Context) {
	executionHistoryId, err := strconv.Atoi(c.Param("executionHistoryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteExecutionhistory(executionHistoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "executionhistory deleted",
	})
}
