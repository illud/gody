package aplication

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	actionsServices "github.com/gody-server/app/actions/domain/services"
	actionsDatabase "github.com/gody-server/app/actions/infraestructure"

	// Replace for dto
	actionsModel "github.com/gody-server/app/actions/domain/models"
)

// Create a actionsDb instance
var actionsDb = actionsDatabase.NewActionsDb()

// Create a Service instance using the actionsDb
var service = actionsServices.NewService(actionsDb)

// Post Actions
// @Summary Post Actions
// @Schemes
// @Description Post Actions
// @Tags Actions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body actionsModel.Actions true "Body to create Actions"
// @Success 200
// @Router /actions [Post]
func CreateActions(c *gin.Context) {
	var actions actionsModel.Actions
	if err := c.ShouldBindJSON(&actions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateActions(actions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "actions created",
	})
}

// Get Actions
// @Summary Get Actions
// @Schemes
// @Description Get Actions
// @Tags Actions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /actions [Get]
func GetActions(c *gin.Context) {
	result, err := service.GetActions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Get Actions
// @Summary Get Actions
// @Schemes
// @Description Get Actions
// @Tags Actions
// @Security BearerAuth
// @Param actionsId path int true "actionsId"
// @Accept json
// @Produce json
// @Success 200
// @Router /actions/{actionsId} [Get]
func GetOneActions(c *gin.Context) {
	actionsId, err := strconv.Atoi(c.Param("actionsId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.GetOneActions(actionsId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Put Actions
// @Summary Put Actions
// @Schemes
// @Description Put Actions
// @Tags Actions
// @Security BearerAuth
// @Param actionsId path int true "actionsId"
// @Accept json
// @Produce json
// @Param Body body actionsModel.Actions true "Body to update Actions"
// @Success 200
// @Router /actions/{actionsId} [Put]
func UpdateActions(c *gin.Context) {
	var actions actionsModel.Actions
	if err := c.ShouldBindJSON(&actions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actionsId, err := strconv.Atoi(c.Param("actionsId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateActions(actionsId, actions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "actions updated",
	})
}

// Delete Actions
// @Summary Delete Actions
// @Schemes
// @Description Delete Actions
// @Tags Actions
// @Security BearerAuth
// @Param actionsId path int true "actionsId"
// @Accept json
// @Produce json
// @Success 200
// @Router /actions/{actionsId} [Delete]
func DeleteActions(c *gin.Context) {
	actionsId, err := strconv.Atoi(c.Param("actionsId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteActions(actionsId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "actions deleted",
	})
}

// Post Actions
// @Summary Post Actions
// @Schemes
// @Description Post Actions
// @Tags Actions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body actionsModel.ActionRun true "Body to create ActionRun"
// @Success 200
// @Router /actions/run [Post]
func Run(c *gin.Context) {
	var actions actionsModel.ActionRun
	if err := c.ShouldBindJSON(&actions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Run(actions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "actions executed successfully",
	})
}
